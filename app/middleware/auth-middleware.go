package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/hiboedi/go-store-backend/app/auth"
	"github.com/hiboedi/go-store-backend/app/exceptions"
	"github.com/hiboedi/go-store-backend/app/helpers"
)

func isPublicRoute(r *http.Request) bool {
	return (r.URL.Path == "/api/login" || r.URL.Path == "/api/signup") && r.Method == "POST"
}

func RedirectSwagger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/swagger") && r.URL.Path == "/api/swagger" {
			http.Redirect(w, r, "/api/swagger/index.html", http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isPublicRoute(r) || strings.HasPrefix(r.URL.Path, "/api/swagger/") {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		if err := auth.VerifyToken(tokenString); err != nil {
			http.Error(w, fmt.Sprintf("Invalid token: %v", err), http.StatusUnauthorized)
			return
		}

		if _, err := r.Cookie(helpers.UserSession); err != nil {
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/api/login", http.StatusFound)
				return
			}
			http.Error(w, "Invalid user cookie", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				exceptions.ErrorHandler(w, r, err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
