package main

import (
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Routes struct {
	ht HandlersTemplates
	hd HandlersData
}

func NewRoutes(ht HandlersTemplates, hd HandlersData) *Routes {
	return &Routes{
		ht: ht,
		hd: hd,
	}
}

func (r *Routes) staticRoutes(mux *http.ServeMux) {
	log.Println("staticRoutes")
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))
}

func (r *Routes) swaggerRoutes(mux *http.ServeMux) {
	log.Println("swaggerRoutes")
	mux.Handle(
		"GET /swagger/",
		loggingMiddleware(httpSwagger.Handler(httpSwagger.URL("/static/swagger.json"))),
	)
}

func (r *Routes) templatesRoutes(mux *http.ServeMux) {
	log.Println("templatesRoutes")
	mux.HandleFunc("GET /", r.ht.home)
	mux.HandleFunc("GET /help", r.ht.help)
	mux.HandleFunc("GET /api", r.ht.api)
	mux.HandleFunc("GET /contact", r.ht.contact)
	mux.HandleFunc("GET /bootstrap", r.ht.bootstrap)
	mux.HandleFunc("GET /clone", r.ht.clone)
	mux.HandleFunc("GET /cmd", r.ht.cmd)
	mux.HandleFunc("GET /config", r.ht.config)
	mux.HandleFunc("GET /console", r.ht.console)
	mux.HandleFunc("GET /convert", r.ht.convert)
	mux.HandleFunc("GET /cp", r.ht.cp)
	mux.HandleFunc("GET /create", r.ht.create)
	mux.HandleFunc("GET /destroy", r.ht.destroy)
	mux.HandleFunc("GET /edit", r.ht.edit)
	mux.HandleFunc("GET /etcupdate", r.ht.etcupdate)
	mux.HandleFunc("GET /export", r.ht.export)
	mux.HandleFunc("GET /htop", r.ht.htop)
	mux.HandleFunc("GET /import", r.ht.imporT)
	mux.HandleFunc("GET /jcp", r.ht.jcp)
	mux.HandleFunc("GET /limits", r.ht.limits)
	mux.HandleFunc("GET /list", r.ht.list)
	mux.HandleFunc("GET /migrate", r.ht.migrate)
	mux.HandleFunc("GET /monitor", r.ht.monitor)
	mux.HandleFunc("GET /mount", r.ht.mount)
	mux.HandleFunc("GET /network", r.ht.network)
	mux.HandleFunc("GET /pkg", r.ht.pkg)
	mux.HandleFunc("GET /rcp", r.ht.rcp)
	mux.HandleFunc("GET /rdr", r.ht.rdr)
	mux.HandleFunc("GET /rename", r.ht.rename)
	mux.HandleFunc("GET /restart", r.ht.restart)
	mux.HandleFunc("GET /service", r.ht.service)
	mux.HandleFunc("GET /setup", r.ht.setup)
	mux.HandleFunc("GET /start", r.ht.start)
	mux.HandleFunc("GET /stop", r.ht.stop)
	mux.HandleFunc("GET /sysrc", r.ht.sysrc)
	mux.HandleFunc("GET /tags", r.ht.tags)
	mux.HandleFunc("GET /template", r.ht.template)
	mux.HandleFunc("GET /top", r.ht.top)
	mux.HandleFunc("GET /umount", r.ht.umount)
	mux.HandleFunc("GET /update", r.ht.update)
	mux.HandleFunc("GET /upgrade", r.ht.upgrade)
	mux.HandleFunc("GET /verify", r.ht.verify)
	mux.HandleFunc("GET /zfs", r.ht.zfs)
}

func (r *Routes) dataRoutes(mux *http.ServeMux) {
	log.Println("dataRoutes")
	mux.HandleFunc("POST /bootstrap", loggingMiddleware(r.hd.bootstrap))
	mux.HandleFunc("POST /clone", loggingMiddleware(r.hd.clone))
	mux.HandleFunc("POST /cmd", loggingMiddleware(r.hd.cmd))
	mux.HandleFunc("POST /config", loggingMiddleware(r.hd.config))
	mux.HandleFunc("POST /console", loggingMiddleware(r.hd.console))
	mux.HandleFunc("POST /convert", loggingMiddleware(r.hd.convert))
	mux.HandleFunc("POST /cp", loggingMiddleware(r.hd.cp))
	mux.HandleFunc("POST /create", loggingMiddleware(r.hd.create))
	mux.HandleFunc("POST /destroy", loggingMiddleware(r.hd.destroy))
	mux.HandleFunc("POST /edit", loggingMiddleware(r.hd.edit))
	mux.HandleFunc("POST /etcupdate", loggingMiddleware(r.hd.etcupdate))
	mux.HandleFunc("POST /export", loggingMiddleware(r.hd.export))
	mux.HandleFunc("POST /htop", loggingMiddleware(r.hd.htop))
	mux.HandleFunc("POST /import", loggingMiddleware(r.hd.imporT))
	mux.HandleFunc("POST /jcp", loggingMiddleware(r.hd.jcp))
	mux.HandleFunc("POST /limits", loggingMiddleware(r.hd.limits))
	mux.HandleFunc("POST /list", loggingMiddleware(r.hd.list))
	mux.HandleFunc("OPTIONS /list", loggingMiddleware(r.hd.list))
	mux.HandleFunc("POST /migrate", loggingMiddleware(r.hd.migrate))
	mux.HandleFunc("POST /monitor", loggingMiddleware(r.hd.monitor))
	mux.HandleFunc("POST /mount", loggingMiddleware(r.hd.mount))
	mux.HandleFunc("POST /network", loggingMiddleware(r.hd.network))
	mux.HandleFunc("POST /pkg", loggingMiddleware(r.hd.pkg))
	mux.HandleFunc("POST /rcp", loggingMiddleware(r.hd.rcp))
	mux.HandleFunc("POST /rdr", loggingMiddleware(r.hd.rdr))
	mux.HandleFunc("POST /rename", loggingMiddleware(r.hd.rename))
	mux.HandleFunc("POST /restart", loggingMiddleware(r.hd.restart))
	mux.HandleFunc("POST /service", loggingMiddleware(r.hd.service))
	mux.HandleFunc("POST /setup", loggingMiddleware(r.hd.setup))
	mux.HandleFunc("POST /start", loggingMiddleware(r.hd.start))
	mux.HandleFunc("POST /stop", loggingMiddleware(r.hd.stop))
	mux.HandleFunc("POST /sysrc", loggingMiddleware(r.hd.sysrc))
	mux.HandleFunc("POST /tags", loggingMiddleware(r.hd.tags))
	mux.HandleFunc("POST /template", loggingMiddleware(r.hd.template))
	mux.HandleFunc("POST /top", loggingMiddleware(r.hd.top))
	mux.HandleFunc("POST /umount", loggingMiddleware(r.hd.umount))
	mux.HandleFunc("POST /update", loggingMiddleware(r.hd.update))
	mux.HandleFunc("POST /upgrade", loggingMiddleware(r.hd.upgrade))
	mux.HandleFunc("POST /verify", loggingMiddleware(r.hd.verify))
	mux.HandleFunc("POST /zfs", loggingMiddleware(r.hd.zfs))
	mux.HandleFunc("GET /node", loggingMiddleware(r.hd.node))
	mux.HandleFunc("OPTIONS /node", loggingMiddleware(r.hd.node))
	mux.HandleFunc("GET /listall", loggingMiddleware(r.hd.listAll))
}
