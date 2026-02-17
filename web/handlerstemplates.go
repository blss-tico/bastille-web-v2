package web

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"bastille-web-v2/bastille"
	"bastille-web-v2/config"
)

type HandlersTemplates struct {
	Bl bastille.Bastille
}

func (ht *HandlersTemplates) login(w http.ResponseWriter, r *http.Request) {
	log.Println("loginHandlersTemplates")
	type SysInfo struct {
		CommandName string
		Ip          string
		Port        string
	}

	host := strings.Split(r.Host, ":")

	data := SysInfo{
		CommandName: "login",
		Ip:          host[0],
		Port:        config.BwPortModel,
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	renderTemplateUtil(w, "login.html", data)
}

func (ht *HandlersTemplates) home(w http.ResponseWriter, r *http.Request) {
	log.Println("homeHandlersTemplates")

	type SysInfo struct {
		Hostname        string `json:"hostname"`
		Arch            string `json:"arch"`
		Platform        string `json:"platform"`
		Osrelease       string `json:"osrelease"`
		Totalmemory     string `json:"totalmemory"`
		BastilleVersion string `json:"bastilleversion"`
		Ip              string `json:"ip"`
		Port            string `json:"port"`
	}

	osinf, _ := bastille.InfoOsUtil()
	posinf := strings.Split(osinf, " ")

	mminf, _ := bastille.MemInfoOsUtil()
	re := regexp.MustCompile(`\d+`)
	pmminf := re.FindAllString(mminf, -1)

	bstv, _ := ht.Bl.BastilleVersion()
	host := strings.Split(r.Host, ":")

	var sysinfo SysInfo
	if len(posinf) > 0 && len(pmminf) > 0 && bstv != "" {
		sysinfo = SysInfo{
			Hostname:        posinf[1],
			Arch:            posinf[len(posinf)-1],
			Platform:        posinf[0],
			Osrelease:       posinf[2],
			Totalmemory:     pmminf[0],
			BastilleVersion: bstv,
			Ip:              host[0],
			Port:            config.BwPortModel,
		}
	}

	config.LoadNodesFile()

	type HomeData struct {
		CommandName string
		Data        config.BastilleModel
		SysInfo     SysInfo
		Nodes       []config.NodesModel
	}

	data := HomeData{
		CommandName: "home",
		Data:        config.BastilleM,
		SysInfo:     sysinfo,
		Nodes:       config.NodesListModel,
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	renderTemplateUtil(w, "home.html", data)
}

func (ht *HandlersTemplates) configuration(w http.ResponseWriter, r *http.Request) {
	log.Println("configurationHandlersTemplates")
	data := templatesModel{CommandName: "configuration", Data: config.BastilleM}
	renderTemplateUtil(w, "configuration.html", data)
}

func (ht *HandlersTemplates) help(w http.ResponseWriter, r *http.Request) {
	log.Println("helpHandlersTemplates")
	data := templatesModel{CommandName: "help", Data: config.BastilleM}
	renderTemplateUtil(w, "help.html", data)
}

func (ht *HandlersTemplates) api(w http.ResponseWriter, r *http.Request) {
	log.Println("apiHandlersTemplates")
	http.Redirect(w, r, "/swagger", http.StatusMovedPermanently)
}

func (ht *HandlersTemplates) contact(w http.ResponseWriter, r *http.Request) {
	log.Println("contactHandlersTemplates")

	type ContactModel struct {
		CommandName string
		Data        config.BastilleModel
		Name        string
		Email       string
		Githubpers  string
		Githubproj  string
	}

	const name = "Bruno Leonardo Tico)"
	const email = "bruno.ccutp@gmail.com"
	const githubpers = "https://github.com/blss-tico"
	const githubproj = "https://github.com/blss-tico/bastille-web"
	data := ContactModel{
		CommandName: "contact",
		Data:        config.BastilleM,
		Name:        name,
		Email:       email,
		Githubpers:  githubpers,
		Githubproj:  githubproj,
	}

	renderTemplateUtil(w, "contact.html", data)
}

func (ht *HandlersTemplates) bootstrap(w http.ResponseWriter, r *http.Request) {
	log.Println("bootstrapHandlersTemplates")
	data := templatesModel{CommandName: "bootstrap", Data: config.BastilleM}
	renderTemplateUtil(w, "bootstrap.html", data)
}

func (ht *HandlersTemplates) clone(w http.ResponseWriter, r *http.Request) {
	log.Println("cloneHandlersTemplates")
	data := templatesModel{CommandName: "clone", Data: config.BastilleM}
	renderTemplateUtil(w, "clone.html", data)
}

func (ht *HandlersTemplates) cmd(w http.ResponseWriter, r *http.Request) {
	log.Println("cmdHandlersTemplates")
	data := templatesModel{CommandName: "cmd", Data: config.BastilleM}
	renderTemplateUtil(w, "cmd.html", data)
}

func (ht *HandlersTemplates) config(w http.ResponseWriter, r *http.Request) {
	log.Println("cmdHandlersTemplates")
	data := templatesModel{CommandName: "config", Data: config.BastilleM}
	renderTemplateUtil(w, "config.html", data)
}

func (ht *HandlersTemplates) console(w http.ResponseWriter, r *http.Request) {
	log.Println("consoleHandlersTemplates")
	data := templatesModel{CommandName: "console", Data: config.BastilleM}
	renderTemplateUtil(w, "console.html", data)
}

func (ht *HandlersTemplates) convert(w http.ResponseWriter, r *http.Request) {
	log.Println("convertHandlersTemplates")
	data := templatesModel{CommandName: "convert", Data: config.BastilleM}
	renderTemplateUtil(w, "convert.html", data)
}

func (ht *HandlersTemplates) cp(w http.ResponseWriter, r *http.Request) {
	log.Println("cpHandlersTemplates")
	data := templatesModel{CommandName: "cp", Data: config.BastilleM}
	renderTemplateUtil(w, "cp.html", data)
}

func (ht *HandlersTemplates) create(w http.ResponseWriter, r *http.Request) {
	log.Println("createHandlersTemplates")
	data := templatesModel{CommandName: "create", Data: config.BastilleM}
	renderTemplateUtil(w, "create.html", data)
}

func (ht *HandlersTemplates) destroy(w http.ResponseWriter, r *http.Request) {
	log.Println("destroyHandlersTemplates")
	data := templatesModel{CommandName: "destroy", Data: config.BastilleM}
	renderTemplateUtil(w, "destroy.html", data)
}

func (ht *HandlersTemplates) edit(w http.ResponseWriter, r *http.Request) {
	log.Println("editHandlersTemplates")
	data := templatesModel{CommandName: "edit", Data: config.BastilleM}
	renderTemplateUtil(w, "edit.html", data)
}

func (ht *HandlersTemplates) etcupdate(w http.ResponseWriter, r *http.Request) {
	log.Println("etcupdateHandlersTemplates")
	data := templatesModel{CommandName: "etcupdate", Data: config.BastilleM}
	renderTemplateUtil(w, "etcupdate.html", data)
}

func (ht *HandlersTemplates) export(w http.ResponseWriter, r *http.Request) {
	log.Println("exportHandlersTemplates")
	data := templatesModel{CommandName: "export", Data: config.BastilleM}
	renderTemplateUtil(w, "export.html", data)
}

func (ht *HandlersTemplates) htop(w http.ResponseWriter, r *http.Request) {
	log.Println("htopHandlersTemplates")
	data := templatesModel{CommandName: "htop", Data: config.BastilleM}
	renderTemplateUtil(w, "htop.html", data)
}

func (ht *HandlersTemplates) imporT(w http.ResponseWriter, r *http.Request) {
	log.Println("importHandlersTemplates")
	data := templatesModel{CommandName: "import", Data: config.BastilleM}
	renderTemplateUtil(w, "import.html", data)
}

func (ht *HandlersTemplates) jcp(w http.ResponseWriter, r *http.Request) {
	log.Println("jcpHandlersTemplates")
	data := templatesModel{CommandName: "jcp", Data: config.BastilleM}
	renderTemplateUtil(w, "jcp.html", data)
}

func (ht *HandlersTemplates) limits(w http.ResponseWriter, r *http.Request) {
	log.Println("limitsHandlersTemplates")
	data := templatesModel{CommandName: "limits", Data: config.BastilleM}
	renderTemplateUtil(w, "limits.html", data)
}

func (ht *HandlersTemplates) list(w http.ResponseWriter, r *http.Request) {
	log.Println("listHandlersTemplates")
	data := templatesModel{CommandName: "list", Data: config.BastilleM}
	renderTemplateUtil(w, "list.html", data)
}

func (ht *HandlersTemplates) migrate(w http.ResponseWriter, r *http.Request) {
	log.Println("migrateHandlersTemplates")
	data := templatesModel{CommandName: "migrate", Data: config.BastilleM}
	renderTemplateUtil(w, "migrate.html", data)
}

func (ht *HandlersTemplates) monitor(w http.ResponseWriter, r *http.Request) {
	log.Println("monitorHandlersTemplates")
	data := templatesModel{CommandName: "monitor", Data: config.BastilleM}
	renderTemplateUtil(w, "monitor.html", data)
}

func (ht *HandlersTemplates) mount(w http.ResponseWriter, r *http.Request) {
	log.Println("mountHandlersTemplates")
	data := templatesModel{CommandName: "mount", Data: config.BastilleM}
	renderTemplateUtil(w, "mount.html", data)
}

func (ht *HandlersTemplates) network(w http.ResponseWriter, r *http.Request) {
	log.Println("networkHandlersTemplates")
	data := templatesModel{CommandName: "network", Data: config.BastilleM}
	renderTemplateUtil(w, "network.html", data)
}

func (ht *HandlersTemplates) pkg(w http.ResponseWriter, r *http.Request) {
	log.Println("pkgHandlersTemplates")
	data := templatesModel{CommandName: "pkg", Data: config.BastilleM}
	renderTemplateUtil(w, "pkg.html", data)
}

func (ht *HandlersTemplates) rcp(w http.ResponseWriter, r *http.Request) {
	log.Println("rcpHandlersTemplates")
	data := templatesModel{CommandName: "rcp", Data: config.BastilleM}
	renderTemplateUtil(w, "rcp.html", data)
}

func (ht *HandlersTemplates) rdr(w http.ResponseWriter, r *http.Request) {
	log.Println("rdrHandlersTemplates")
	data := templatesModel{CommandName: "rdr", Data: config.BastilleM}
	renderTemplateUtil(w, "rdr.html", data)
}

func (ht *HandlersTemplates) rename(w http.ResponseWriter, r *http.Request) {
	log.Println("renameHandlersTemplates")
	data := templatesModel{CommandName: "rename", Data: config.BastilleM}
	renderTemplateUtil(w, "rename.html", data)
}

func (ht *HandlersTemplates) restart(w http.ResponseWriter, r *http.Request) {
	log.Println("restartHandlersTemplates")
	data := templatesModel{CommandName: "restart", Data: config.BastilleM}
	renderTemplateUtil(w, "restart.html", data)
}

func (ht *HandlersTemplates) service(w http.ResponseWriter, r *http.Request) {
	log.Println("serviceHandlersTemplates")
	data := templatesModel{CommandName: "service", Data: config.BastilleM}
	renderTemplateUtil(w, "service.html", data)
}

func (ht *HandlersTemplates) setup(w http.ResponseWriter, r *http.Request) {
	log.Println("setupHandlersTemplates")
	data := templatesModel{CommandName: "setup", Data: config.BastilleM}
	renderTemplateUtil(w, "setup.html", data)
}

func (ht *HandlersTemplates) start(w http.ResponseWriter, r *http.Request) {
	log.Println("startHandlersTemplates")
	data := templatesModel{CommandName: "start", Data: config.BastilleM}
	renderTemplateUtil(w, "start.html", data)
}

func (ht *HandlersTemplates) stop(w http.ResponseWriter, r *http.Request) {
	log.Println("stopHandlersTemplates")
	data := templatesModel{CommandName: "stop", Data: config.BastilleM}
	renderTemplateUtil(w, "stop.html", data)
}

func (ht *HandlersTemplates) sysrc(w http.ResponseWriter, r *http.Request) {
	log.Println("sysrcHandlersTemplates")
	data := templatesModel{CommandName: "sysrc", Data: config.BastilleM}
	renderTemplateUtil(w, "sysrc.html", data)
}

func (ht *HandlersTemplates) tags(w http.ResponseWriter, r *http.Request) {
	log.Println("tagsHandlersTemplates")
	data := templatesModel{CommandName: "tags", Data: config.BastilleM}
	renderTemplateUtil(w, "tags.html", data)
}

func (ht *HandlersTemplates) template(w http.ResponseWriter, r *http.Request) {
	log.Println("templateHandlersTemplates")
	data := templatesModel{CommandName: "template", Data: config.BastilleM}
	renderTemplateUtil(w, "template.html", data)
}

func (ht *HandlersTemplates) top(w http.ResponseWriter, r *http.Request) {
	log.Println("topHandlersTemplates")
	data := templatesModel{CommandName: "top", Data: config.BastilleM}
	renderTemplateUtil(w, "top.html", data)
}

func (ht *HandlersTemplates) umount(w http.ResponseWriter, r *http.Request) {
	log.Println("umountHandlersTemplates")
	data := templatesModel{CommandName: "umount", Data: config.BastilleM}
	renderTemplateUtil(w, "umount.html", data)
}

func (ht *HandlersTemplates) update(w http.ResponseWriter, r *http.Request) {
	log.Println("updateHandlersTemplates")
	data := templatesModel{CommandName: "update", Data: config.BastilleM}
	renderTemplateUtil(w, "update.html", data)
}

func (ht *HandlersTemplates) upgrade(w http.ResponseWriter, r *http.Request) {
	log.Println("upgradeHandlersTemplates")
	data := templatesModel{CommandName: "upgrade", Data: config.BastilleM}
	renderTemplateUtil(w, "upgrade.html", data)
}

func (ht *HandlersTemplates) verify(w http.ResponseWriter, r *http.Request) {
	log.Println("verifyHandlersTemplates")
	data := templatesModel{CommandName: "verify", Data: config.BastilleM}
	renderTemplateUtil(w, "verify.html", data)
}

func (ht *HandlersTemplates) zfs(w http.ResponseWriter, r *http.Request) {
	log.Println("zfsHandlersTemplates")
	data := templatesModel{CommandName: "zfs", Data: config.BastilleM}
	renderTemplateUtil(w, "zfs.html", data)
}
