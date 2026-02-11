package users

import (
	"bastille-web-v2/config"

	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func loggingMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionId := r.Header.Get("X-Request-ID")
		start := time.Now()
		log.Printf("[users]: %s %s %s %v %s", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start), sessionId)
		f(w, r)
	}
}

func SessionAuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("SessionAuthMiddleware")

		cookie, err := r.Cookie("bw-actk")
		if err != nil {
			if err == http.ErrNoCookie {
				log.Println("No bw-actk cookie found redirect")
				http.Redirect(w, r, "/refreshtpt", http.StatusUnauthorized)
				return
			} else {
				log.Println("Error reading bw-actk cookie redirect")
				http.Redirect(w, r, "/refreshtpt", http.StatusUnauthorized)
				return
			}
		} else {
			log.Println("Cookie bw-actk found")
			ctx := context.WithValue(r.Context(), "bw-actk", cookie.Value)
			r = r.WithContext(ctx)
		}

		f(w, r)
	}
}

func ApiAuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("ApiAuthMiddleware")

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid Authorization format", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &claimsModel{}
		pToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JwtKeyModel, nil
		})
		if err != nil || !pToken.Valid {
			http.Error(w, "Invalid Authorization token", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "username", claims.Username)
		f(w, r.WithContext(ctx))
	}
}
