package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/bharharya/openstack-compute-service/utils"
)

// AuthMiddleware validates JWT tokens in incoming requests
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // No "Bearer " prefix
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Store user ID in request context
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserIDFromContext extracts userID from request context
func GetUserIDFromContext(r *http.Request) (uint, error) {
	userID, ok := r.Context().Value("userID").(float64) // JWT stores numbers as float64
	if !ok {
		return 0, errors.New("userID not found in context")
	}
	return uint(userID), nil
}
