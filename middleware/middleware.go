package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify the token (this is a placeholder, you'll need to replace with actual JWT verification)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Normally, you'd validate the signing algorithm and return the secret key
			return []byte("your-256-bit-secret"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
