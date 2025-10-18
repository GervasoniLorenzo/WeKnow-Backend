package middleware

import "net/http"

// Richiede che il proxy abbia aggiunto X-Admin: true (Basic Auth superata)
func AdminOnly() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-Admin") != "true" {
				http.Error(w, "admin gateway required", http.StatusForbidden)
				return
			}
			// opzionale: anche role == "admin"
			// if c := ClaimsFromCtx(r.Context()); c.Role != "admin" {
			// 	http.Error(w, "admin role required", http.StatusForbidden)
			// 	return
			// }
			next.ServeHTTP(w, r)
		})
	}
}
