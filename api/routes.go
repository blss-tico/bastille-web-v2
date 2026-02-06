package api

import (
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
		apiAuthMiddleware(loggingMiddleware(httpSwagger.Handler(httpSwagger.URL("/static/swagger.json")))),
	)
}

func (r *Routes) DataRoutes(mux *http.ServeMux) {
	log.Println("dataRoutes")
	mux.HandleFunc("POST /bootstrap", apiAuthMiddleware(loggingMiddleware(r.Hd.bootstrap)))
	mux.HandleFunc("POST /clone", loggingMiddleware(r.Hd.clone))
	mux.HandleFunc("POST /cmd", loggingMiddleware(r.Hd.cmd))
	mux.HandleFunc("POST /config", loggingMiddleware(r.Hd.config))
	mux.HandleFunc("POST /console", loggingMiddleware(r.Hd.console))
	mux.HandleFunc("POST /convert", loggingMiddleware(r.Hd.convert))
	mux.HandleFunc("POST /cp", loggingMiddleware(r.Hd.cp))
	mux.HandleFunc("POST /create", loggingMiddleware(r.Hd.create))
	mux.HandleFunc("POST /destroy", loggingMiddleware(r.Hd.destroy))
	mux.HandleFunc("POST /edit", loggingMiddleware(r.Hd.edit))
	mux.HandleFunc("POST /etcupdate", loggingMiddleware(r.Hd.etcupdate))
	mux.HandleFunc("POST /export", loggingMiddleware(r.Hd.export))
	mux.HandleFunc("POST /htop", loggingMiddleware(r.Hd.htop))
	mux.HandleFunc("POST /import", loggingMiddleware(r.Hd.imporT))
	mux.HandleFunc("POST /jcp", loggingMiddleware(r.Hd.jcp))
	mux.HandleFunc("POST /limits", loggingMiddleware(r.Hd.limits))
	mux.HandleFunc("POST /list", loggingMiddleware(r.Hd.list))
	mux.HandleFunc("OPTIONS /list", loggingMiddleware(r.Hd.list))
	mux.HandleFunc("POST /migrate", loggingMiddleware(r.Hd.migrate))
	mux.HandleFunc("POST /monitor", loggingMiddleware(r.Hd.monitor))
	mux.HandleFunc("POST /mount", loggingMiddleware(r.Hd.mount))
	mux.HandleFunc("POST /network", loggingMiddleware(r.Hd.network))
	mux.HandleFunc("POST /pkg", loggingMiddleware(r.Hd.pkg))
	mux.HandleFunc("POST /rcp", loggingMiddleware(r.Hd.rcp))
	mux.HandleFunc("POST /rdr", loggingMiddleware(r.Hd.rdr))
	mux.HandleFunc("POST /rename", loggingMiddleware(r.Hd.rename))
	mux.HandleFunc("POST /restart", loggingMiddleware(r.Hd.restart))
	mux.HandleFunc("POST /service", loggingMiddleware(r.Hd.service))
	mux.HandleFunc("POST /setup", loggingMiddleware(r.Hd.setup))
	mux.HandleFunc("POST /start", loggingMiddleware(r.Hd.start))
	mux.HandleFunc("POST /stop", loggingMiddleware(r.Hd.stop))
	mux.HandleFunc("POST /sysrc", loggingMiddleware(r.Hd.sysrc))
	mux.HandleFunc("POST /tags", loggingMiddleware(r.Hd.tags))
	mux.HandleFunc("POST /template", loggingMiddleware(r.Hd.template))
	mux.HandleFunc("POST /top", loggingMiddleware(r.Hd.top))
	mux.HandleFunc("POST /umount", loggingMiddleware(r.Hd.umount))
	mux.HandleFunc("POST /update", loggingMiddleware(r.Hd.update))
	mux.HandleFunc("POST /upgrade", loggingMiddleware(r.Hd.upgrade))
	mux.HandleFunc("POST /verify", loggingMiddleware(r.Hd.verify))
	mux.HandleFunc("POST /zfs", loggingMiddleware(r.Hd.zfs))
	mux.HandleFunc("GET /node", loggingMiddleware(r.Hd.node))
	mux.HandleFunc("OPTIONS /node", loggingMiddleware(r.Hd.node))
	mux.HandleFunc("GET /listall", loggingMiddleware(r.Hd.listAll))
}
