package middleware

import (
	"context"
	"net/http"
	"strings"

	"book-management/utils"
)

type contextKey string

const (
	userIDContextKey   contextKey = "user_id"
	userRoleContextKey contextKey = "user_role"
)

// Auth middleware - verifies JWT token
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			Error(w, http.StatusUnauthorized, "Authorization header required")
			return
		}

		// Support "Bearer <token>" format
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			Error(w, http.StatusUnauthorized, "Invalid token format")
			return
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			Error(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		// Add user info to request context
		ctx := context.WithValue(r.Context(), userIDContextKey, claims.UserID)
		ctx = context.WithValue(ctx, userRoleContextKey, claims.Role)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// GetUserID - Get user ID from context
func GetUserID(r *http.Request) int64 {
	if id, ok := r.Context().Value(userIDContextKey).(int64); ok {
		return id
	}
	return 0
}

// GetUserRole - Get user role from context
func GetUserRole(r *http.Request) string {
	if role, ok := r.Context().Value(userRoleContextKey).(string); ok {
		return role
	}
	return "user"
}