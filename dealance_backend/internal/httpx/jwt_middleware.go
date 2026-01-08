package httpx

import (
	"context"
	"net/http"
	"strings"

	"dealance.co/backend/internal/crypto"
)

func JWTMiddleware(jwtManager *crypto.JWTManager) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			auth := r.Header.Get("Authorization")
			if auth == "" {
				http.Error(w, "missing authorization", http.StatusUnauthorized)
				return
			}

			parts := strings.SplitN(auth, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid authorization format", http.StatusUnauthorized)
				return
			}

			claims, err := jwtManager.VerifyToken(parts[1])
			if err != nil {
				http.Error(w, "invalid or expired token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), ContextUserID, claims.UserID)
			ctx = context.WithValue(ctx, ContextUserRole, claims.Role)
			ctx = context.WithValue(ctx, ContextKYCLevel, claims.KYCLevel)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
