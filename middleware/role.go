package middleware

import "net/http"

// Role middleware - checks if user has required role
func Role(allowedRoles ...string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			userRole := GetUserRole(r)

			// Allow if user has any of the allowed roles
			for _, allowed := range allowedRoles {
				if userRole == allowed {
					next.ServeHTTP(w, r)
					return
				}
			}

			// Special case: admin has access to everything
			if userRole == "admin" {
				next.ServeHTTP(w, r)
				return
			}

			Error(w, http.StatusForbidden, "Insufficient permissions. Required role: "+allowedRoles[0])
		}
	}
}