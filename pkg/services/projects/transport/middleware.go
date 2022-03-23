package transport

import (
	"net/http"
	"strings"
)

// corsAccessHeader
func corsAccessHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
		w.Header().Set("Access-Control-Allow-Headers", strings.Join([]string{"Content-Type", "Origin", "Accept"}, ","))
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}

// jsonContentHeader
func jsonContentHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
