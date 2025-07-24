package helpers

import (
	"net/http"
	"strings"
)

// Wrap any http.Handler, add CORS headers, and handle OPTIONS preflight.
func CORSMiddleware(allowedOrigins []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// If you want to restrict to specific origins:
		if len(allowedOrigins) > 0 {
			for _, o := range allowedOrigins {
				if o == "*" || strings.EqualFold(o, origin) {
					w.Header().Set("Access-Control-Allow-Origin", o)
					goto setCommon
				}
			}
			// origin not allowed
			http.Error(w, "CORS origin denied", http.StatusForbidden)
			return
		}

		// If you want to allow * everything:
		w.Header().Set("Access-Control-Allow-Origin", "*")

	setCommon:
		// common CORS headers
		w.Header().Set("Vary", "Origin")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// short‑circuit preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
