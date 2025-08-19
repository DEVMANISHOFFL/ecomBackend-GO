package auth

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const ClaimsKey = contextKey("claims")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]
		claims, err := ParseJWT(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if claims.UserID == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Attach claims to context
		ctx := context.WithValue(r.Context(), ClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetClaimsFromContext(r *http.Request) *Claims {
	val := r.Context().Value(ClaimsKey)
	if claims, ok := val.(*Claims); ok {
		return claims
	}
	return nil
}

func RoleMiddleware(allowedRoles ...Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims := GetClaimsFromContext(r)
			if claims == nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			for _, role := range allowedRoles {
				if claims.Role == role {
					next.ServeHTTP(w, r)
					return
				}
			}

			http.Error(w, "Forbidden", http.StatusForbidden)
		})
	}
}
