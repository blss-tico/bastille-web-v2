package nodes

import (
	"bastille-web-v2/users"

	"log"
	"net/http"
)

type Routes struct {
	Hn HandlersNodes
}

func NewRoutes(hn HandlersNodes) *Routes {
	return &Routes{
		Hn: hn,
	}
}

func (r *Routes) NodesRoutes(mux *http.ServeMux) {
	log.Println("userRoutes")

	mux.HandleFunc("POST /nodes", loggingMiddleware(users.SessionAuthMiddleware(r.Hn.createNodes)))
	mux.HandleFunc("GET /nodes", loggingMiddleware(users.SessionAuthMiddleware(r.Hn.getNodes)))
	mux.HandleFunc("PUT /nodes/{nodename}", loggingMiddleware(users.SessionAuthMiddleware(r.Hn.updateNodes)))
	mux.HandleFunc("DELETE /nodes/{nodename}", loggingMiddleware(users.SessionAuthMiddleware(r.Hn.deleteNodes)))
}
