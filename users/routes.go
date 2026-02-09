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

	routes := map[string]http.HandlerFunc{
		"POST /logout":       r.Hu.logout,
		"POST /refresh":      r.Hu.refresh,
		"GET /users":         r.Hu.getUsers,
		"PUT /users/{id}":    r.Hu.updateUser,
		"DELETE /users/{id}": r.Hu.deleteUser,
	}

	for path, handler := range routes {
		cmd := loggingMiddleware(SessionAuthMiddleware(http.HandlerFunc(handler)))
		http.Handle(path, cmd)
	}
}
