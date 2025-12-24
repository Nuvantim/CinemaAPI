package config

import (
	"log"
	"net/http"
	"time"
)

func LoggingRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			userAgent := r.UserAgent()
			endpoint := r.URL.Path

			log.Printf(
				"[%s] %s %v | User-Agent: %s",
				r.Method,
				endpoint,
				time.Since(start),
				userAgent,
			)
		}()
		next.ServeHTTP(w, r)
	})
}
