package api

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logga tutte le richieste HTTP
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("➡️ %s %s", r.Method, r.URL.Path)

		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)

		log.Printf("⬅️ %s %s (%v)", r.Method, r.URL.Path, time.Since(start))
	})
}


func RecoveryMiddleware(next http.Handler) http.Handler {

	return nil
}


func JSONMiddleware(next http.Handler) http.Handler {

	return nil
}


func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ChainMiddleware applica più middleware
func ChainMiddleware(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}














