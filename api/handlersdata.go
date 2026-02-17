package api

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"

	"bastille-web-v2/bastille"

	"github.com/go-playground/validator/v10"
)

type HandlersData struct {
	Bl bastille.Bastille
}

// health
// @Summary health api check
// @Description Route to check if bastille-web api is running.
// @Tags health
// @Produce  application/json
// @Success 200 {object} string
// @Router /health [get]
func (hd *HandlersData) health(w http.ResponseWriter, r *http.Request) {
	log.Println("healthHandler")

	type HealthResponse struct {
		Status string `json:"status"`
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := HealthResponse{Status: "bastille-web api running"}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// bootstrap
// @Summary bootstrap command
// @Description The bootstrap sub-command is used to download and extract releases and templates for use with Bastille containers. A valid release is needed before containers can be created. Templates are optional but are managed in the same manner.
// @Tags bootstrap
// @Accept  json
// @Produce  text/plain
// @Param  bootstrap  body	bootstrapModel  true  "bootstrap"
// @Success 200 {object} string
// @Router /bootstrap [post]
func (hd *HandlersData) bootstrap(w http.ResponseWriter, r *http.Request) {
	log.Println("bootstrapHandler")

	var data bootstrapModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Bootstrap(data.Options, data.ReleaseTemplate, data.UpdateArch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// clone
// @Summary clone command
// @Description Clone/duplicate an existing jail to a new jail.
// @Tags clone
// @Accept  json
// @Produce  text/plain
// @Param  clone  body	cloneModel  true  "clone"
// @Success 200 {object} string
// @Router /clone [post]
func (hd *HandlersData) clone(w http.ResponseWriter, r *http.Request) {
	log.Println("cloneHandler")

	var data cloneModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Clone(data.Options, data.Target, data.Newname, data.Ip)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// cmd
// @Summary cmd command
// @Description Execute commands inside targeted jail(s).
// @Tags cmd
// @Accept  json
// @Produce  text/plain
// @Param  cmd  body	cmdModel  true  "cmd"
// @Success 200 {object} string
// @Router /cmd [post]
func (hd *HandlersData) cmd(w http.ResponseWriter, r *http.Request) {
	log.Println("cmdHandler")

	var data cmdModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	commands := strings.Split(data.Command, " ")
	if len(commands) == 0 {
		http.Error(w, "command not found", http.StatusBadRequest)
		return
	}

	result, err := hd.Bl.Cmd(data.Options, data.Target, commands)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// config
// @Summary config command
// @Description Get, set or remove properties from targeted jail(s).
// @Tags config
// @Accept  json
// @Produce  text/plain
// @Param  config  body	configModel  true  "config"
// @Success 200 {object} string
// @Router /config [post]
func (hd *HandlersData) config(w http.ResponseWriter, r *http.Request) {
	log.Println("configHandler")

	var data configModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Config(data.Options, data.Target, data.Action, data.Property, data.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// console
// @Summary console command
// @Description Launch a login shell into the jail. Default is password- less root login.
// @Tags console
// @Accept  json
// @Produce  text/plain
// @Param  console  body	consoleModel  true  "config"
// @Success 200 {object} string
// @Router /console [post]
func (hd *HandlersData) console(w http.ResponseWriter, r *http.Request) {
	log.Println("consoleHandler")

	var data consoleModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Console(data.Options, data.Target, data.User)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type ConsoleResponse struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := ConsoleResponse{Host: r.Host, Port: result}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// convert
// @Summary convert command
// @Description Convert a thin jail to a thick jail. Convert a thick jail to a custom release.
// @Tags convert
// @Accept  json
// @Produce  text/plain
// @Param  convert  body	convertModel  true  "config"
// @Success 200 {object} string
// @Router /convert [post]
func (hd *HandlersData) convert(w http.ResponseWriter, r *http.Request) {
	log.Println("convertHandler")

	var data convertModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Convert(data.Options, data.Target, data.Release)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// cp
// @Summary cp command
// @Description Copy files from host to jail(s).
// @Tags cp
// @Accept  json
// @Produce  text/plain
// @Param  cp  body	cpModel  true  "cp"
// @Success 200 {object} string
// @Router /cp [post]
func (hd *HandlersData) cp(w http.ResponseWriter, r *http.Request) {
	log.Println("cpHandler")

	var data cpModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Cp(data.Options, data.Target, data.Hostpath, data.Jailpath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// create
// @Summary create command
// @Description Create a jail uning any availaBle bootstrapped release. To create a jail, simply provide a name, bootstrapped release, and IP address.
// @Tags create
// @Accept  json
// @Produce  text/plain
// @Param  create  body	createModel  true  "create"
// @Success 200 {object} string
// @Router /create [post]
func (hd *HandlersData) create(w http.ResponseWriter, r *http.Request) {
	log.Println("createHandler")

	var data createModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Create(data.Options, data.Name, data.Release, data.Ip, data.Iface, data.Gtwip, data.Ipip, data.Value, data.Vlanid, data.Zfsoptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// destroy
// @Summary destroy command
// @Description Destroy jails or releases.
// @Tags destroy
// @Accept  json
// @Produce  text/plain
// @Param  destroy  body	destroyModel  true  "destroy"
// @Success 200 {object} string
// @Router /destroy [post]
func (hd *HandlersData) destroy(w http.ResponseWriter, r *http.Request) {
	log.Println("destroyHandler")

	var data destroyModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Destroy(data.Options, data.JailRelease)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// edit
// @Summary edit command
// @Description Edit jail config files.
// @Tags edit
// @Accept  json
// @Produce  text/plain
// @Param  edit  body	editModel  true  "edit"
// @Success 200 {object} string
// @Router /edit [post]
func (hd *HandlersData) edit(w http.ResponseWriter, r *http.Request) {
	log.Println("editHandler")

	var data editModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Edit(data.Options, data.Target, data.File)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type EditResponse struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := EditResponse{Host: r.Host, Port: result}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// etcupdate
// @Summary etcupdate command
// @Description This command will update the contents of /etc inside a jail. It should be run after a jail upgrade.
// @Tags etcupdate
// @Accept  json
// @Produce  text/plain
// @Param  etcupdate  body	etcupdateModel  true  "etcupdate"
// @Success 200 {object} string
// @Router /etcupdate [post]
func (hd *HandlersData) etcupdate(w http.ResponseWriter, r *http.Request) {
	log.Println("etcupdateHandler")

	var data etcupdateModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Etcupdate(data.Options, data.Bootstraptarget, data.Action, data.Release)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// export
// @Summary export command
// @Description Exporting a container creates an archive or image that can be sent to a different machine to be imported later. These exported archives can be used as container backups.
// @Tags export
// @Accept  json
// @Produce  text/plain
// @Param  export  body	exportModel  true  "export"
// @Success 200 {object} string
// @Router /export [post]
func (hd *HandlersData) export(w http.ResponseWriter, r *http.Request) {
	log.Println("exportHandler")

	var data exportModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	options := strings.Split(data.Options, " ")
	if len(options) == 0 {
		http.Error(w, "options error", http.StatusBadRequest)
		return
	}

	result, err := hd.Bl.Export(options, data.Target, data.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// htop
// @Summary htop command
// @Description This command runs htop in the targeted jail. Requires htop to be installed in the jail.
// @Tags htop
// @Accept  json
// @Produce  text/plain
// @Param  htop  body	htopModel  true  "htop"
// @Success 200 {object} string
// @Router /htop [post]
func (hd *HandlersData) htop(w http.ResponseWriter, r *http.Request) {
	log.Println("htopHandler")

	var data htopModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Htop(data.Options, data.Target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type HtopResponse struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := HtopResponse{Host: r.Host, Port: result}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// import
// @Summary import command
// @Description Import a jail backup image or archive.
// @Tags import
// @Accept  json
// @Produce  text/plain
// @Param  import  body	importModel  true  "import"
// @Success 200 {object} string
// @Router /import [post]
func (hd *HandlersData) imporT(w http.ResponseWriter, r *http.Request) {
	log.Println("importHandler")

	var data importModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.ImporT(data.Options, data.File, data.Release)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// jcp
// @Summary jcp command
// @Description Copy files from jail to jail(s).
// @Tags jcp
// @Accept  json
// @Produce  text/plain
// @Param  jcp  body	jcpModel  true  "jcp"
// @Success 200 {object} string
// @Router /jcp [post]
func (hd *HandlersData) jcp(w http.ResponseWriter, r *http.Request) {
	log.Println("jcpHandler")

	var data jcpModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Jcp(data.Options, data.Sourcejail, data.Jailpath, data.Destjail, data.Jailpath2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// limits
// @Summary limits command
// @Description Set resourse limits for targeted jail(s).
// @Tags limits
// @Accept  json
// @Produce  text/plain
// @Param  limits  body	limitsModel  true  "limits"
// @Success 200 {object} string
// @Router /limits [post]
func (hd *HandlersData) limits(w http.ResponseWriter, r *http.Request) {
	log.Println("limitsHandler")

	var data limitsModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Limits(data.Options, data.Target, data.Action, data.Option, data.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// list
// @Summary list command
// @Description List jails, ports, releases, templates, logs, limits, exports and imports and much more managed by bastille.
// @Tags list
// @Accept  json
// @Produce  text/plain
// @Param  list  body	listModel  true  "list"
// @Success 200 {object} string
// @Router /list [post]
func (hd *HandlersData) list(w http.ResponseWriter, r *http.Request) {
	log.Println("listHandler")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var data listModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.List(data.Options, data.Action)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func (hd *HandlersData) listAll(w http.ResponseWriter, r *http.Request) {
	log.Println("listAllHandler")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	result, err := hd.Bl.List("-j", "all")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// migrate
// @Summary migrate command
// @Description The migrate sub-command allows migrating the targeted jail(s) to another remote system.
// @Tags migrate
// @Accept  json
// @Produce  text/plain
// @Param  migrate  body	migrateModel  true  "migrate"
// @Success 200 {object} string
// @Router /migrate [post]
func (hd *HandlersData) migrate(w http.ResponseWriter, r *http.Request) {
	log.Println("migrateHandler")

	var data migrateModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Migrate(data.Options, data.Target, data.Remote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// monitor
// @Summary monitor command
// @Description The monitor sub-command adds, removes, lists and enables/disables monitoring for container services.
// @Tags monitor
// @Accept  json
// @Produce  text/plain
// @Param  monitor body	monitorModel true  "monitor"
// @Success 200 {object} string
// @Router /monitor [post]
func (hd *HandlersData) monitor(w http.ResponseWriter, r *http.Request) {
	log.Println("monitorHandler")

	var data monitorModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Monitor(data.Options, data.Target, data.Action, data.Service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// mount
// @Summary mount command
// @Description To mount storage within the container use bastille mount.
// @Tags mount
// @Accept  json
// @Produce  text/plain
// @Param  mount  body	mountModel  true  "mount"
// @Success 200 {object} string
// @Router /mount [post]
func (hd *HandlersData) mount(w http.ResponseWriter, r *http.Request) {
	log.Println("mountHandler")

	var data mountModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Mount(data.Options, data.Target, data.Hostpath, data.Jailpath, data.Filesystemtype, data.Option, data.Dump, data.Passnumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// network
// @Summary network command
// @Description Add or remove interfaces to existing jails.
// @Tags network
// @Accept  json
// @Produce  text/plain
// @Param  network  body	networkModel  true  "network"
// @Success 200 {object} string
// @Router /network [post]
func (hd *HandlersData) network(w http.ResponseWriter, r *http.Request) {
	log.Println("networkHandler")

	var data networkModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Network(data.Options, data.Target, data.Action, data.Iface, data.Ip, data.Vlanid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// pkg
// @Summary pkg command
// @Description Manage binary packages inside jails.
// @Tags pkg
// @Accept  json
// @Produce  text/plain
// @Param  pkg  body	pkgModel  true  "pkg"
// @Success 200 {object} string
// @Router /pkg [post]
func (hd *HandlersData) pkg(w http.ResponseWriter, r *http.Request) {
	log.Println("pkgHandler")

	var data pkgModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	arg := strings.Split(data.Arg, " ")
	if len(arg) == 0 {
		http.Error(w, "args is not found", http.StatusBadRequest)
		return
	}

	result, err := hd.Bl.Pkg(data.Options, data.Target, arg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// rcp
// @Summary rcp command
// @Description This command allows copying files from jail to host.
// @Tags rcp
// @Accept  json
// @Produce  text/plain
// @Param  rcp  body	rcpModel  true  "rcp"
// @Success 200 {object} string
// @Router /rcp [post]
func (hd *HandlersData) rcp(w http.ResponseWriter, r *http.Request) {
	log.Println("rcpHandler")

	var data rcpModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Rcp(data.Options, data.Target, data.Jailpath, data.Hostpath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// rdr
// @Summary rdr command
// @Description Bastille rdr allows you to configure dynamic rdr rules for your containers without modifying pf.conf.
// @Tags rdr
// @Accept  json
// @Produce  text/plain
// @Param  rdr  body	rdrModel  true  "rdr"
// @Success 200 {object} string
// @Router /rdr [post]
func (hd *HandlersData) rdr(w http.ResponseWriter, r *http.Request) {
	log.Println("rdrHandler")

	var data rdrModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Rdr(data.Options, data.Odestination, data.Ointerface, data.Osource, data.Otype, data.Target, data.Action, data.Hostport, data.Jailport, data.Log, data.Logopts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// rename
// @Summary rename command
// @Description Rename a jail.
// @Tags rename
// @Accept  json
// @Produce  text/plain
// @Param  rename  body	renameModel  true  "rename"
// @Success 200 {object} string
// @Router /rename [post]
func (hd *HandlersData) rename(w http.ResponseWriter, r *http.Request) {
	log.Println("renameHandler")

	var data renameModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Rename(data.Options, data.Target, data.Newname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// restart
// @Summary restart command
// @Description Restart jail(s).
// @Tags restart
// @Accept  json
// @Produce  text/plain
// @Param  restart  body	restartModel  true  "restart"
// @Success 200 {object} string
// @Router /restart [post]
func (hd *HandlersData) restart(w http.ResponseWriter, r *http.Request) {
	log.Println("restartHandler")

	var data restartModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Restart(data.Options, data.Target, data.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// service
// @Summary service command
// @Description The service sub-command allows for managing services within jails. This allows you to start, stop, restart, and otherwise interact with services running inside the jail(s).
// @Tags service
// @Accept  json
// @Produce  text/plain
// @Param  service  body	serviceModel  true  "service"
// @Success 200 {object} string
// @Router /service [post]
func (hd *HandlersData) service(w http.ResponseWriter, r *http.Request) {
	log.Println("serviceHandler")

	var data serviceModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Service(data.Options, data.Target, data.Servicename, data.Args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// setup
// @Summary setup command
// @Description The setup sub-command attempts to automatically configure a host system for Bastille jails. This allows you to configure networking, firewall, storage, vnet and bridge options for a Bastille host with one command.
// @Tags setup
// @Accept  json
// @Produce  text/plain
// @Param  setup  body	setupModel  true  "setup"
// @Success 200 {object} string
// @Router /setup [post]
func (hd *HandlersData) setup(w http.ResponseWriter, r *http.Request) {
	log.Println("setupHandler")

	var data setupModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Setup(data.Options, data.Action)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// start
// @Summary start command
// @Description Start jail(s).
// @Tags start
// @Accept  json
// @Produce  text/plain
// @Param  start  body	startModel  true  "start"
// @Success 200 {object} string
// @Router /start [post]
func (hd *HandlersData) start(w http.ResponseWriter, r *http.Request) {
	log.Println("startHandler")

	var data startModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Start(data.Options, data.Target, data.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// stop
// @Summary stop command
// @Description Stop jail(s).
// @Tags stop
// @Accept  json
// @Produce  text/plain
// @Param  stop  body	stopModel  true  "stop"
// @Success 200 {object} string
// @Router /stop [post]
func (hd *HandlersData) stop(w http.ResponseWriter, r *http.Request) {
	log.Println("stopHandler")

	var data stopModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Stop(data.Options, data.Target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// sysrc
// @Summary sysrc command
// @Description The sysrc sub-command allows for safely editing system configuration files. In jail terms, this allows us to toggle on/off services and options at startup.
// @Tags sysrc
// @Accept  json
// @Produce  text/plain
// @Param  sysrc  body	sysrcModel  true  "sysrc"
// @Success 200 {object} string
// @Router /sysrc [post]
func (hd *HandlersData) sysrc(w http.ResponseWriter, r *http.Request) {
	log.Println("sysrcHandler")

	var data sysrcModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Sysrc(data.Options, data.Target, data.Args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// tags
// @Summary tags command
// @Description The tags sub-command adds, removes or lists arbitrary tags on your jail(s).
// @Tags tags
// @Accept  json
// @Produce  text/plain
// @Param  tags  body	tagsModel  true  "tags"
// @Success 200 {object} string
// @Router /tags [post]
func (hd *HandlersData) tags(w http.ResponseWriter, r *http.Request) {
	log.Println("tagsHandler")

	var data tagsModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Tags(data.Options, data.Target, data.Action, data.Tgs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// template
// @Summary template command
// @Description Apply file templates to targeted container(s).
// @Tags template
// @Accept  json
// @Produce  text/plain
// @Param  template  body	templateModel  true  "template"
// @Success 200 {object} string
// @Router /template [post]
func (hd *HandlersData) template(w http.ResponseWriter, r *http.Request) {
	log.Println("templateHandler")

	var data templateModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Template(data.Options, data.Target, data.Action, data.Template)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// top
// @Summary top command
// @Description This command runs top in the targeted jail.
// @Tags top
// @Accept  json
// @Produce  text/plain
// @Param  top  body	topModel  true  "top"
// @Success 200 {object} string
// @Router /top [post]
func (hd *HandlersData) top(w http.ResponseWriter, r *http.Request) {
	log.Println("topHandler")

	var data topModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Top(data.Options, data.Target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type TopResponse struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := TopResponse{Host: r.Host, Port: result}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// umount
// @Summary umount command
// @Description Unmount storage from jail(s).
// @Tags umount
// @Accept  json
// @Produce  text/plain
// @Param  umount  body	umountModel  true  "umount"
// @Success 200 {object} string
// @Router /umount [post]
func (hd *HandlersData) umount(w http.ResponseWriter, r *http.Request) {
	log.Println("umountHandler")

	var data umountModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Umount(data.Options, data.Target, data.Jailpath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// update
// @Summary update command
// @Description The update command targets a release or a thick jail. Because thin jails are based on a release, when the release is updated all the thin jails are automatically updated as well.
// @Tags update
// @Accept  json
// @Produce  text/plain
// @Param  update  body	updateModel  true  "update"
// @Success 200 {object} string
// @Router /update [post]
func (hd *HandlersData) update(w http.ResponseWriter, r *http.Request) {
	log.Println("updateHandler")

	var data updateModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Update(data.Options, data.Target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// upgrade
// @Summary upgrade command
// @Description The upgrade command targets a thick or thin jail. Thin jails will be updated by changing the release mount point that it is based on. Thick jails will be upgraded normally.
// @Tags upgrade
// @Accept  json
// @Produce  text/plain
// @Param  upgrade  body	upgradeModel  true  "upgrade"
// @Success 200 {object} string
// @Router /upgrade [post]
func (hd *HandlersData) upgrade(w http.ResponseWriter, r *http.Request) {
	log.Println("upgradeHandler")

	var data upgradeModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Upgrade(data.Options, data.Target, data.Action)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// verify
// @Summary verify command
// @Description This command scans a bootstrapped release or template and validates that everything looks in order. This is not a 100% comprehensive check, but it compares the release or template against a “known good” index.
// @Tags verify
// @Accept  json
// @Produce  text/plain
// @Param  verify  body	verifyModel  true  "verify"
// @Success 200 {object} string
// @Router /verify [post]
func (hd *HandlersData) verify(w http.ResponseWriter, r *http.Request) {
	log.Println("verifyHandler")

	var data verifyModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Verify(data.Options, data.Action)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// zfs
// @Summary zfs command
// @Description Manage ZFS properties, create, destroy and rollback snapshots, jail and unjail datasets (ZFS only), and check ZFS usage for targeted jail(s).
// @Tags zfs
// @Accept  json
// @Produce  text/plain
// @Param  zfs  body	zfsModel  true  "zfs"
// @Success 200 {object} string
// @Router /zfs [post]
func (hd *HandlersData) zfs(w http.ResponseWriter, r *http.Request) {
	log.Println("zfsHandler")

	var data zfsModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := hd.Bl.Zfs(data.Options, data.Target, data.Action, data.Tag, data.Key, data.Value, data.Pooldataset, data.Jailpath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// node
// @Summary node information
// @Description Get information about bastille-web host node.
// @Tags node
// @Accept  json
// @Produce  application/json
// @Success 200 {object} string
// @Router /node [get]
func (hd *HandlersData) node(w http.ResponseWriter, r *http.Request) {
	log.Println("nodeHandler")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte{})
		return
	}

	type SysInfo struct {
		Hostname        string `json:"hostname"`
		Arch            string `json:"arch"`
		Platform        string `json:"platform"`
		Osrelease       string `json:"osrelease"`
		Totalmemory     string `json:"totalmemory"`
		BastilleVersion string `json:"bastilleversion"`
	}

	osinf, _ := bastille.InfoOsUtil()
	posinf := strings.Split(osinf, " ")

	mminf, _ := bastille.MemInfoOsUtil()
	re := regexp.MustCompile(`\d+`)
	pmminf := re.FindAllString(mminf, -1)

	bstv, _ := hd.Bl.BastilleVersion()

	var sysinfo SysInfo
	if len(posinf) > 0 && len(pmminf) > 0 && bstv != "" {
		sysinfo = SysInfo{
			Hostname:        posinf[1],
			Arch:            posinf[len(posinf)-1],
			Platform:        posinf[0],
			Osrelease:       posinf[2],
			Totalmemory:     pmminf[0],
			BastilleVersion: bstv,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(sysinfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
