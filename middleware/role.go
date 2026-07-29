package middleware

import (
	"net/http"

	"book-management/utils"
)

func Role(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole := GetUserRole(r)

			for _, allowed := range allowedRoles {
				if userRole == allowed {
					next.ServeHTTP(w, r)
					return
				}
			}

			if userRole == "admin" {
				next.ServeHTTP(w, r)
				return
			}

			utils.Error(w, http.StatusForbidden, "Insufficient permissions")
		})
	}
}