package api

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
		log.Printf("%s %s %s %v %s", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start), sessionId)
		f(w, r)
	}
}

func apiAuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "username", claims.Username)
		f(w, r.WithContext(ctx))
	}
}
