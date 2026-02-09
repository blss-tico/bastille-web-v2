package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Routes struct {
	Ht HandlersTemplates
}

func NewRoutes(ht HandlersTemplates) *Routes {
	return &Routes{
		Ht: ht,
	}
}

func (r *Routes) StaticRoutes(mux *http.ServeMux) {
	log.Println("staticRoutes")
	fs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))
}

func (rt *Routes) serveLogin(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")

	htmlFile, err := os.ReadFile("./web/templates/login.html")
	if err != nil {
		http.Error(w, "Can't find login.html", http.StatusInternalServerError)
		log.Println("Can't find login.html")
		return
	}

	fmt.Fprint(w, string(htmlFile))
}

func (r *Routes) TemplatesRoutes(mux *http.ServeMux) {
	log.Println("templatesRoutes")
	mux.HandleFunc("GET /", r.Ht.login)
	mux.HandleFunc("GET /home", r.Ht.home)
	mux.HandleFunc("GET /help", r.Ht.help)
	mux.HandleFunc("GET /api", r.Ht.api)
	mux.HandleFunc("GET /contact", r.Ht.contact)
	mux.HandleFunc("GET /bootstrap", r.Ht.bootstrap)
	mux.HandleFunc("GET /clone", r.Ht.clone)
	mux.HandleFunc("GET /cmd", r.Ht.cmd)
	mux.HandleFunc("GET /config", r.Ht.config)
	mux.HandleFunc("GET /console", r.Ht.console)
	mux.HandleFunc("GET /convert", r.Ht.convert)
	mux.HandleFunc("GET /cp", r.Ht.cp)
	mux.HandleFunc("GET /create", r.Ht.create)
	mux.HandleFunc("GET /destroy", r.Ht.destroy)
	mux.HandleFunc("GET /edit", r.Ht.edit)
	mux.HandleFunc("GET /etcupdate", r.Ht.etcupdate)
	mux.HandleFunc("GET /export", r.Ht.export)
	mux.HandleFunc("GET /htop", r.Ht.htop)
	mux.HandleFunc("GET /import", r.Ht.imporT)
	mux.HandleFunc("GET /jcp", r.Ht.jcp)
	mux.HandleFunc("GET /limits", r.Ht.limits)
	mux.HandleFunc("GET /list", r.Ht.list)
	mux.HandleFunc("GET /migrate", r.Ht.migrate)
	mux.HandleFunc("GET /monitor", r.Ht.monitor)
	mux.HandleFunc("GET /mount", r.Ht.mount)
	mux.HandleFunc("GET /network", r.Ht.network)
	mux.HandleFunc("GET /pkg", r.Ht.pkg)
	mux.HandleFunc("GET /rcp", r.Ht.rcp)
	mux.HandleFunc("GET /rdr", r.Ht.rdr)
	mux.HandleFunc("GET /rename", r.Ht.rename)
	mux.HandleFunc("GET /restart", r.Ht.restart)
	mux.HandleFunc("GET /service", r.Ht.service)
	mux.HandleFunc("GET /setup", r.Ht.setup)
	mux.HandleFunc("GET /start", r.Ht.start)
	mux.HandleFunc("GET /stop", r.Ht.stop)
	mux.HandleFunc("GET /sysrc", r.Ht.sysrc)
	mux.HandleFunc("GET /tags", r.Ht.tags)
	mux.HandleFunc("GET /template", r.Ht.template)
	mux.HandleFunc("GET /top", r.Ht.top)
	mux.HandleFunc("GET /umount", r.Ht.umount)
	mux.HandleFunc("GET /update", r.Ht.update)
	mux.HandleFunc("GET /upgrade", r.Ht.upgrade)
	mux.HandleFunc("GET /verify", r.Ht.verify)
	mux.HandleFunc("GET /zfs", r.Ht.zfs)
}
