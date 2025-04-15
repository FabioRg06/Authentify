package middleware

import (
	"log"
	"net/http"
)

// Middleware para logging
func LogRequests(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("📩 %s request to %s", r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}
