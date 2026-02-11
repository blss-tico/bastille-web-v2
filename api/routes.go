package api

import (
	"bastille-web-v2/users"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Routes struct {
	Hd HandlersData
}

func NewRoutes(hd HandlersData) *Routes {
	return &Routes{
		Hd: hd,
	}
}

func (r *Routes) SwaggerRoutes(mux *http.ServeMux) {
	log.Println("swaggerRoutes")
	mux.Handle(
		"GET /swagger/",
		users.ApiAuthMiddleware(loggingMiddleware(httpSwagger.Handler(httpSwagger.URL("/static/swagger.json")))),
	)
}

func (r *Routes) DataRoutes(mux *http.ServeMux) {
	log.Println("dataRoutes")

	mux.HandleFunc("GET /health", loggingMiddleware(r.Hd.bootstrap))
	mux.HandleFunc("POST /bootstrap", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.bootstrap)))
	mux.HandleFunc("POST /clone", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.clone)))
	mux.HandleFunc("POST /cmd", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.cmd)))
	mux.HandleFunc("POST /config", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.config)))
	mux.HandleFunc("POST /console", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.console)))
	mux.HandleFunc("POST /convert", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.convert)))
	mux.HandleFunc("POST /cp", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.cp)))
	mux.HandleFunc("POST /create", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.create)))
	mux.HandleFunc("POST /destroy", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.destroy)))
	mux.HandleFunc("POST /edit", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.edit)))
	mux.HandleFunc("POST /etcupdate", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.etcupdate)))
	mux.HandleFunc("POST /export", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.export)))
	mux.HandleFunc("POST /htop", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.htop)))
	mux.HandleFunc("POST /import", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.imporT)))
	mux.HandleFunc("POST /jcp", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.jcp)))
	mux.HandleFunc("POST /limits", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.limits)))
	mux.HandleFunc("POST /list", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.list)))
	mux.HandleFunc("OPTIONS /list", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.list)))
	mux.HandleFunc("POST /migrate", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.migrate)))
	mux.HandleFunc("POST /monitor", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.monitor)))
	mux.HandleFunc("POST /mount", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.mount)))
	mux.HandleFunc("POST /network", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.network)))
	mux.HandleFunc("POST /pkg", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.pkg)))
	mux.HandleFunc("POST /rcp", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.rcp)))
	mux.HandleFunc("POST /rdr", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.rdr)))
	mux.HandleFunc("POST /rename", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.rename)))
	mux.HandleFunc("POST /restart", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.restart)))
	mux.HandleFunc("POST /service", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.service)))
	mux.HandleFunc("POST /setup", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.setup)))
	mux.HandleFunc("POST /start", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.start)))
	mux.HandleFunc("POST /stop", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.stop)))
	mux.HandleFunc("POST /sysrc", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.sysrc)))
	mux.HandleFunc("POST /tags", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.tags)))
	mux.HandleFunc("POST /template", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.template)))
	mux.HandleFunc("POST /top", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.top)))
	mux.HandleFunc("POST /umount", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.umount)))
	mux.HandleFunc("POST /update", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.update)))
	mux.HandleFunc("POST /upgrade", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.upgrade)))
	mux.HandleFunc("POST /verify", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.verify)))
	mux.HandleFunc("POST /zfs", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.zfs)))
	mux.HandleFunc("GET /node", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.node)))
	mux.HandleFunc("OPTIONS /node", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.node)))
	mux.HandleFunc("GET /listall", loggingMiddleware(users.ApiAuthMiddleware(r.Hd.listAll)))
}
