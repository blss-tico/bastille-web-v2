package web

import (
	"context"
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

func sessionAuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("access_token")
		if err != nil {
			if err == http.ErrNoCookie {
				log.Println("No access_token cookie found")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			} else {
				log.Println("Error reading cookie:", err)
			}
		} else {
			log.Println("Found access_token cookie:", cookie.Value)
			ctx := context.WithValue(r.Context(), "access_token", cookie.Value)
			r = r.WithContext(ctx)
		}

		f(w, r)
	}
}
