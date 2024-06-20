package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/hiboedi/go-store-backend/app/helpers"
)

// Middleware struct
type AuthMiddleware struct {
	Handler http.Handler
}

// Membuat middleware baru
func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

// Fungsi untuk memverifikasi token
func VerifyToken(token string) error {
	// Implementasikan logika untuk memverifikasi token
	// Ini adalah contoh, sesuaikan dengan logika verifikasi yang sesuai
	return nil
}

// Implementasi middleware ServeHTTP
func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Dapatkan token dari header Authorization
	tokenString := r.Header.Get("Authorization")

	// Jika ini adalah permintaan login atau signup, lewati autentikasi
	if (r.URL.Path == "/api/login" || r.URL.Path == "/api/signup") && r.Method == "POST" {
		middleware.Handler.ServeHTTP(w, r)
		return
	}

	// Pastikan header Authorization tersedia
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}

	// Menghapus prefix "Bearer " dari token
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Memeriksa apakah token valid
	if err := VerifyToken(tokenString); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, err.Error())
		return
	}

	// Memeriksa cookie user
	if _, err := r.Cookie(helpers.UserSession); err != nil {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/api/login", http.StatusFound)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid user cookie")
		return
	}

	// Jika semua pengecekan valid, lanjutkan pemrosesan permintaan
	middleware.Handler.ServeHTTP(w, r)
}

// Middleware untuk memulihkan dari panic
func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Tangani kesalahan
				fmt.Fprintf(w, "An error occurred: %v", err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
