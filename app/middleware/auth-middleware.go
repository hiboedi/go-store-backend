package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/hiboedi/go-store-backend/app/auth"
	"github.com/hiboedi/go-store-backend/app/helpers"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func isPublicRoute(r *http.Request) bool {
	return (r.URL.Path == "/api/login" || r.URL.Path == "/api/signup") && r.Method == "POST"
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if isPublicRoute(r) {
		middleware.Handler.ServeHTTP(w, r)
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

	middleware.Handler.ServeHTTP(w, r)
}
