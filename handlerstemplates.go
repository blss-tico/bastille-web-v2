package main

import (
	"log"
	"net/http"
	"regexp"
	"strings"
)

type HandlersTemplates struct {
	bl Bastille
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

	osinf, _ := infoOsUtil()
	posinf := strings.Split(osinf, " ")

	mminf, _ := memInfoOsUtil()
	re := regexp.MustCompile(`\d+`)
	pmminf := re.FindAllString(mminf, -1)

	bstv, _ := ht.bl.bastilleVersion()

	var sysinfo SysInfo
	if len(posinf) > 0 && len(pmminf) > 0 && bstv != "" {
		sysinfo = SysInfo{
			Hostname:        posinf[1],
			Arch:            posinf[len(posinf)-1],
			Platform:        posinf[0],
			Osrelease:       posinf[2],
			Totalmemory:     pmminf[0],
			BastilleVersion: bstv,
			Ip:              addrModel,
			Port:            portModel,
		}
	}

	/*
		type ListAllJails struct {
			Action  string `json:"action"`
			Options string `json:"options"`
		}

		allJails := ListAllJails{
			Action:  "all",
			Options: "-j",
		}

		jsonData, err := json.Marshal(allJails)
		if err != nil {
			log.Fatalf("Error marshalling JSON: %v", err)
		}

		url := "http://" + addrModel + ":" + portModel + "/list"
		result, err := http.Post(
			url,
			"application/json",
			bytes.NewBuffer(jsonData),
		)
		if err != nil {
			log.Fatalf("Error making POST request: %v", err)
		}
		defer result.Body.Close()

		if result.StatusCode != http.StatusOK {
			http.Error(w, errors.New("bad status code").Error(), result.StatusCode)
			return
		}

		/*
			result, err := ht.bl.list("-j", "all")
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}


		bodyBytes, err := io.ReadAll(result.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}

		type Response struct {
			Msg string `json:"msg"`
		}

		var resp Response
		if err := json.Unmarshal([]byte(bodyBytes), &resp); err != nil {
			log.Fatalf("Error unmarshaling response: %v", err)
		}

		type List struct {
			Jid       string `json:"jid"`
			Boot      string `json:"boot"`
			Prio      string `json:"prio"`
			State     string `json:"state"`
			Ip        string `json:"ip address"`
			Published string `json:"published ports"`
			Hostname  string `json:"hostname"`
			Release   string `json:"release"`
			Path      string `json:"path"`
		}

		var list []List
		if err := json.Unmarshal([]byte(resp.Msg), &list); err != nil {
			log.Fatalf("Error unmarshaling JSON: %v", err)
		}
	*/

	type HomeData struct {
		CommandName string
		Data        bastilleModel
		//JailsData   []List
		SysInfo SysInfo
		Addr    string
		Nodes   []nodesModel
	}

	data := HomeData{
		CommandName: "home",
		Data:        bastille,
		//JailsData:   list,
		SysInfo: sysinfo,
		Addr:    addrModel,
		Nodes:   nodesListModel,
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	renderTemplateUtil(w, "home.html", data)
}

func (ht *HandlersTemplates) help(w http.ResponseWriter, r *http.Request) {
	log.Println("helpHandlersTemplates")
	data := templatesModel{CommandName: "help", Data: bastille}
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
		Data        bastilleModel
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
		Data:        bastille,
		Name:        name,
		Email:       email,
		Githubpers:  githubpers,
		Githubproj:  githubproj,
	}

	renderTemplateUtil(w, "contact.html", data)
}

func (ht *HandlersTemplates) bootstrap(w http.ResponseWriter, r *http.Request) {
	log.Println("bootstrapHandlersTemplates")
	data := templatesModel{CommandName: "bootstrap", Data: bastille}
	renderTemplateUtil(w, "bootstrap.html", data)
}

func (ht *HandlersTemplates) clone(w http.ResponseWriter, r *http.Request) {
	log.Println("cloneHandlersTemplates")
	data := templatesModel{CommandName: "clone", Data: bastille}
	renderTemplateUtil(w, "clone.html", data)
}

func (ht *HandlersTemplates) cmd(w http.ResponseWriter, r *http.Request) {
	log.Println("cmdHandlersTemplates")
	data := templatesModel{CommandName: "cmd", Data: bastille}
	renderTemplateUtil(w, "cmd.html", data)
}

func (ht *HandlersTemplates) config(w http.ResponseWriter, r *http.Request) {
	log.Println("cmdHandlersTemplates")
	data := templatesModel{CommandName: "config", Data: bastille}
	renderTemplateUtil(w, "config.html", data)
}

func (ht *HandlersTemplates) console(w http.ResponseWriter, r *http.Request) {
	log.Println("consoleHandlersTemplates")
	data := templatesModel{CommandName: "console", Data: bastille}
	renderTemplateUtil(w, "console.html", data)
}

func (ht *HandlersTemplates) convert(w http.ResponseWriter, r *http.Request) {
	log.Println("convertHandlersTemplates")
	data := templatesModel{CommandName: "convert", Data: bastille}
	renderTemplateUtil(w, "convert.html", data)
}

func (ht *HandlersTemplates) cp(w http.ResponseWriter, r *http.Request) {
	log.Println("cpHandlersTemplates")
	data := templatesModel{CommandName: "cp", Data: bastille}
	renderTemplateUtil(w, "cp.html", data)
}

func (ht *HandlersTemplates) create(w http.ResponseWriter, r *http.Request) {
	log.Println("createHandlersTemplates")
	data := templatesModel{CommandName: "create", Data: bastille}
	renderTemplateUtil(w, "create.html", data)
}

func (ht *HandlersTemplates) destroy(w http.ResponseWriter, r *http.Request) {
	log.Println("destroyHandlersTemplates")
	data := templatesModel{CommandName: "destroy", Data: bastille}
	renderTemplateUtil(w, "destroy.html", data)
}

func (ht *HandlersTemplates) edit(w http.ResponseWriter, r *http.Request) {
	log.Println("editHandlersTemplates")
	data := templatesModel{CommandName: "edit", Data: bastille}
	renderTemplateUtil(w, "edit.html", data)
}

func (ht *HandlersTemplates) etcupdate(w http.ResponseWriter, r *http.Request) {
	log.Println("etcupdateHandlersTemplates")
	data := templatesModel{CommandName: "etcupdate", Data: bastille}
	renderTemplateUtil(w, "etcupdate.html", data)
}

func (ht *HandlersTemplates) export(w http.ResponseWriter, r *http.Request) {
	log.Println("exportHandlersTemplates")
	data := templatesModel{CommandName: "export", Data: bastille}
	renderTemplateUtil(w, "export.html", data)
}

func (ht *HandlersTemplates) htop(w http.ResponseWriter, r *http.Request) {
	log.Println("htopHandlersTemplates")
	data := templatesModel{CommandName: "htop", Data: bastille}
	renderTemplateUtil(w, "htop.html", data)
}

func (ht *HandlersTemplates) imporT(w http.ResponseWriter, r *http.Request) {
	log.Println("importHandlersTemplates")
	data := templatesModel{CommandName: "import", Data: bastille}
	renderTemplateUtil(w, "import.html", data)
}

func (ht *HandlersTemplates) jcp(w http.ResponseWriter, r *http.Request) {
	log.Println("jcpHandlersTemplates")
	data := templatesModel{CommandName: "jcp", Data: bastille}
	renderTemplateUtil(w, "jcp.html", data)
}

func (ht *HandlersTemplates) limits(w http.ResponseWriter, r *http.Request) {
	log.Println("limitsHandlersTemplates")
	data := templatesModel{CommandName: "limits", Data: bastille}
	renderTemplateUtil(w, "limits.html", data)
}

func (ht *HandlersTemplates) list(w http.ResponseWriter, r *http.Request) {
	log.Println("listHandlersTemplates")
	data := templatesModel{CommandName: "list", Data: bastille}
	renderTemplateUtil(w, "list.html", data)
}

func (ht *HandlersTemplates) migrate(w http.ResponseWriter, r *http.Request) {
	log.Println("migrateHandlersTemplates")
	data := templatesModel{CommandName: "migrate", Data: bastille}
	renderTemplateUtil(w, "migrate.html", data)
}

func (ht *HandlersTemplates) monitor(w http.ResponseWriter, r *http.Request) {
	log.Println("monitorHandlersTemplates")
	data := templatesModel{CommandName: "monitor", Data: bastille}
	renderTemplateUtil(w, "monitor.html", data)
}

func (ht *HandlersTemplates) mount(w http.ResponseWriter, r *http.Request) {
	log.Println("mountHandlersTemplates")
	data := templatesModel{CommandName: "mount", Data: bastille}
	renderTemplateUtil(w, "mount.html", data)
}

func (ht *HandlersTemplates) network(w http.ResponseWriter, r *http.Request) {
	log.Println("networkHandlersTemplates")
	data := templatesModel{CommandName: "network", Data: bastille}
	renderTemplateUtil(w, "network.html", data)
}

func (ht *HandlersTemplates) pkg(w http.ResponseWriter, r *http.Request) {
	log.Println("pkgHandlersTemplates")
	data := templatesModel{CommandName: "pkg", Data: bastille}
	renderTemplateUtil(w, "pkg.html", data)
}

func (ht *HandlersTemplates) rcp(w http.ResponseWriter, r *http.Request) {
	log.Println("rcpHandlersTemplates")
	data := templatesModel{CommandName: "rcp", Data: bastille}
	renderTemplateUtil(w, "rcp.html", data)
}

func (ht *HandlersTemplates) rdr(w http.ResponseWriter, r *http.Request) {
	log.Println("rdrHandlersTemplates")
	data := templatesModel{CommandName: "rdr", Data: bastille}
	renderTemplateUtil(w, "rdr.html", data)
}

func (ht *HandlersTemplates) rename(w http.ResponseWriter, r *http.Request) {
	log.Println("renameHandlersTemplates")
	data := templatesModel{CommandName: "rename", Data: bastille}
	renderTemplateUtil(w, "rename.html", data)
}

func (ht *HandlersTemplates) restart(w http.ResponseWriter, r *http.Request) {
	log.Println("restartHandlersTemplates")
	data := templatesModel{CommandName: "restart", Data: bastille}
	renderTemplateUtil(w, "restart.html", data)
}

func (ht *HandlersTemplates) service(w http.ResponseWriter, r *http.Request) {
	log.Println("serviceHandlersTemplates")
	data := templatesModel{CommandName: "service", Data: bastille}
	renderTemplateUtil(w, "service.html", data)
}

func (ht *HandlersTemplates) setup(w http.ResponseWriter, r *http.Request) {
	log.Println("setupHandlersTemplates")
	data := templatesModel{CommandName: "setup", Data: bastille}
	renderTemplateUtil(w, "setup.html", data)
}

func (ht *HandlersTemplates) start(w http.ResponseWriter, r *http.Request) {
	log.Println("startHandlersTemplates")
	data := templatesModel{CommandName: "start", Data: bastille}
	renderTemplateUtil(w, "start.html", data)
}

func (ht *HandlersTemplates) stop(w http.ResponseWriter, r *http.Request) {
	log.Println("stopHandlersTemplates")
	data := templatesModel{CommandName: "stop", Data: bastille}
	renderTemplateUtil(w, "stop.html", data)
}

func (ht *HandlersTemplates) sysrc(w http.ResponseWriter, r *http.Request) {
	log.Println("sysrcHandlersTemplates")
	data := templatesModel{CommandName: "sysrc", Data: bastille}
	renderTemplateUtil(w, "sysrc.html", data)
}

func (ht *HandlersTemplates) tags(w http.ResponseWriter, r *http.Request) {
	log.Println("tagsHandlersTemplates")
	data := templatesModel{CommandName: "tags", Data: bastille}
	renderTemplateUtil(w, "tags.html", data)
}

func (ht *HandlersTemplates) template(w http.ResponseWriter, r *http.Request) {
	log.Println("templateHandlersTemplates")
	data := templatesModel{CommandName: "template", Data: bastille}
	renderTemplateUtil(w, "template.html", data)
}

func (ht *HandlersTemplates) top(w http.ResponseWriter, r *http.Request) {
	log.Println("topHandlersTemplates")
	data := templatesModel{CommandName: "top", Data: bastille}
	renderTemplateUtil(w, "top.html", data)
}

func (ht *HandlersTemplates) umount(w http.ResponseWriter, r *http.Request) {
	log.Println("umountHandlersTemplates")
	data := templatesModel{CommandName: "umount", Data: bastille}
	renderTemplateUtil(w, "umount.html", data)
}

func (ht *HandlersTemplates) update(w http.ResponseWriter, r *http.Request) {
	log.Println("updateHandlersTemplates")
	data := templatesModel{CommandName: "update", Data: bastille}
	renderTemplateUtil(w, "update.html", data)
}

func (ht *HandlersTemplates) upgrade(w http.ResponseWriter, r *http.Request) {
	log.Println("upgradeHandlersTemplates")
	data := templatesModel{CommandName: "upgrade", Data: bastille}
	renderTemplateUtil(w, "upgrade.html", data)
}

func (ht *HandlersTemplates) verify(w http.ResponseWriter, r *http.Request) {
	log.Println("verifyHandlersTemplates")
	data := templatesModel{CommandName: "verify", Data: bastille}
	renderTemplateUtil(w, "verify.html", data)
}

func (ht *HandlersTemplates) zfs(w http.ResponseWriter, r *http.Request) {
	log.Println("zfsHandlersTemplates")
	data := templatesModel{CommandName: "zfs", Data: bastille}
	renderTemplateUtil(w, "zfs.html", data)
}
