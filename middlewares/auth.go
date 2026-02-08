package middlewares

import (
	"context"
	"net/http"
	"strings"

	"url_shortener/utils"
)

type contextKey string

const userEmailKey contextKey = "userEmail"

// AuthMiddleware validates JWT token
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// Expect: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		email, err := utils.ValidateJWT(parts[1])
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// attach email to request context
		ctx := context.WithValue(r.Context(), userEmailKey, email)
		next(w, r.WithContext(ctx))
	}
}

// GetUserEmailFromContext extracts email from context
func GetUserEmailFromContext(r *http.Request) string {
	email, _ := r.Context().Value(userEmailKey).(string)
	return email
}
