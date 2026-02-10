package web

import (
	"bastille-web-v2/users"
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

	mux.HandleFunc("GET /", loggingMiddleware(r.Ht.login))
	mux.HandleFunc("GET /home", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.home)))
	mux.HandleFunc("GET /configuration", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.configuration)))
	mux.HandleFunc("GET /help", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.help)))
	mux.HandleFunc("GET /api", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.api)))
	mux.HandleFunc("GET /contact", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.contact)))
	mux.HandleFunc("GET /bootstrap", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.bootstrap)))
	mux.HandleFunc("GET /clone", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.clone)))
	mux.HandleFunc("GET /cmd", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.cmd)))
	mux.HandleFunc("GET /config", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.config)))
	mux.HandleFunc("GET /console", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.console)))
	mux.HandleFunc("GET /convert", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.convert)))
	mux.HandleFunc("GET /cp", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.cp)))
	mux.HandleFunc("GET /create", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.create)))
	mux.HandleFunc("GET /destroy", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.destroy)))
	mux.HandleFunc("GET /edit", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.edit)))
	mux.HandleFunc("GET /etcupdate", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.etcupdate)))
	mux.HandleFunc("GET /export", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.export)))
	mux.HandleFunc("GET /htop", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.htop)))
	mux.HandleFunc("GET /import", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.imporT)))
	mux.HandleFunc("GET /jcp", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.jcp)))
	mux.HandleFunc("GET /limits", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.limits)))
	mux.HandleFunc("GET /list", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.list)))
	mux.HandleFunc("GET /migrate", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.migrate)))
	mux.HandleFunc("GET /monitor", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.monitor)))
	mux.HandleFunc("GET /mount", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.mount)))
	mux.HandleFunc("GET /network", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.network)))
	mux.HandleFunc("GET /pkg", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.pkg)))
	mux.HandleFunc("GET /rcp", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.rcp)))
	mux.HandleFunc("GET /rdr", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.rdr)))
	mux.HandleFunc("GET /rename", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.rename)))
	mux.HandleFunc("GET /restart", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.restart)))
	mux.HandleFunc("GET /service", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.service)))
	mux.HandleFunc("GET /setup", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.setup)))
	mux.HandleFunc("GET /start", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.start)))
	mux.HandleFunc("GET /stop", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.stop)))
	mux.HandleFunc("GET /sysrc", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.sysrc)))
	mux.HandleFunc("GET /tags", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.tags)))
	mux.HandleFunc("GET /template", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.template)))
	mux.HandleFunc("GET /top", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.top)))
	mux.HandleFunc("GET /umount", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.umount)))
	mux.HandleFunc("GET /update", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.update)))
	mux.HandleFunc("GET /upgrade", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.upgrade)))
	mux.HandleFunc("GET /verify", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.verify)))
	mux.HandleFunc("GET /zfs", loggingMiddleware(users.SessionAuthMiddleware(r.Ht.zfs)))
}
