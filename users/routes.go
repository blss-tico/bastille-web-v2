package users

import (
	"log"
	"net/http"
)

type Routes struct {
	Hu HandlersUser
}

func NewRoutes(hu HandlersUser) *Routes {
	return &Routes{
		Hu: hu,
	}
}

func (r *Routes) UserRoutes(mux *http.ServeMux) {
	log.Println("userRoutes")

	mux.HandleFunc("POST /register", loggingMiddleware(r.Hu.register))
	mux.HandleFunc("POST /login", loggingMiddleware(r.Hu.login))
	mux.HandleFunc("POST /refresh", loggingMiddleware(r.Hu.refreshTkApi))
	mux.HandleFunc("POST /refreshtpt", loggingMiddleware(r.Hu.refreshCkTpt))
	mux.HandleFunc("POST /logout", loggingMiddleware(SessionAuthMiddleware(r.Hu.logout)))
	mux.HandleFunc("GET /users", loggingMiddleware(SessionAuthMiddleware(r.Hu.getUsers)))
	mux.HandleFunc("PUT /users/{username}", loggingMiddleware(SessionAuthMiddleware(r.Hu.updateUser)))
	mux.HandleFunc("DELETE /users/{username}", loggingMiddleware(SessionAuthMiddleware(r.Hu.deleteUser)))
}
