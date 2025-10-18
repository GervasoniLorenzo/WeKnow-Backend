package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type ctxKey string

var userKey ctxKey = "userClaims"

type Claims struct {
	Sub   string
	Email string
	Role  string
}

func ClaimsFromCtx(ctx context.Context) Claims {
	if v, ok := ctx.Value(userKey).(Claims); ok {
		return v
	}
	return Claims{}
}

func AppJWT(secret []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Leggiamo prima dal cookie
			var tokenStr string
			if c, err := r.Cookie("app_session"); err == nil && c.Value != "" {
				tokenStr = c.Value
			} else {
				// Fallback: Authorization: Bearer <jwt>
				auth := r.Header.Get("Authorization")
				if strings.HasPrefix(auth, "Bearer ") {
					tokenStr = strings.TrimPrefix(auth, "Bearer ")
				}
			}
			if tokenStr == "" {
				http.Error(w, "missing token", http.StatusUnauthorized)
				return
			}

			parsed, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected alg")
				}
				return secret, nil
			})
			if err != nil || !parsed.Valid {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			mc, ok := parsed.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "invalid claims", http.StatusUnauthorized)
				return
			}

			claims := Claims{
				Sub:   fmt.Sprintf("%v", mc["sub"]),
				Email: fmt.Sprintf("%v", mc["email"]),
				Role:  fmt.Sprintf("%v", mc["role"]),
			}
			ctx := context.WithValue(r.Context(), userKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
