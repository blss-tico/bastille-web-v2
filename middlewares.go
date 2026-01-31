package main

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionId := r.Header.Get("X-Request-ID")
		start := time.Now()
		log.Printf("%s %s %s %v %s", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start), sessionId)
		f(w, r)
	}
}
