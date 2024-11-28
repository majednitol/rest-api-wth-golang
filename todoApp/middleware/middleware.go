package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the incoming HTTP request with method and URL
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

// RecoverMiddleware handles panics and returns a 500 error
func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				log.Printf("Recovered from panic: %v", err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
