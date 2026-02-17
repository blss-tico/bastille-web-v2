package api

import "github.com/dgrijalva/jwt-go"

type claimsModel struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type bootstrapModel struct {
	Options         string `json:"options,omitempty" example:"-x" format:"string"`
	ReleaseTemplate string `json:"release|template" example:"14.3-RELEASE" format:"string" validate:"required"`
	UpdateArch      string `json:"update|arch,omitempty" example:"--i386" format:"string"`
}

type cloneModel struct {
	Options string `json:"options,omitempty" example:"-a|-l|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Newname string `json:"new_name" example:"new_jail_name" format:"string" validate:"required"`
	Ip      string `json:"ip" example:"n.n.n.n" format:"string" validate:"ip"`
}

type cmdModel struct {
	Options string `json:"options,omitempty" example:"-a|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Command string `json:"command" example:"ls -l" format:"string" validate:"required"`
}

type configModel struct {
	Options  string `json:"options,omitempty" example:"-x" format:"string"`
	Target   string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Action   string `json:"action" example:"get|set" format:"string" validate:"required"`
	Property string `json:"property" example:"ip4.addr" format:"string" validate:"required"`
	Value    string `json:"value,omitempty" example:"depends on property" format:"string"`
}

type consoleModel struct {
	Options string `json:"options,omitempty" example:"-a|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
	User    string `json:"user,omitempty" example:"root" format:"string"`
}

type convertModel struct {
	Options string `json:"options,omitempty" example:"-a|-y|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Release string `json:"release,omitempty" example:"myrelease" format:"string"`
}

type cpModel struct {
	Options  string `json:"options,omitempty" example:"-q|-x" format:"string"`
	Target   string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Hostpath string `json:"host_path" example:"/host/path" format:"string" validate:"filepath"`
	Jailpath string `json:"jail_path" example:"/jail/path" format:"string" validate:"filepath"`
}

type createModel struct {
	Options    string `json:"options,omitempty" example:"-B|-C|-D|-E|-g|-L|-M|-n|--no-validate|--no-boot|-p|-T|-V|-v|-x|-Z" format:"string"`
	Name       string `json:"name" example:"jail_name" format:"string" validate:"required"`
	Release    string `json:"release" example:"14.3-RELEASE" format:"string" validate:"required"`
	Ip         string `json:"ip" example:"n.n.n.n" format:"string" validate:"omitempty,ip|cidr"`
	Iface      string `json:"iface,omitempty" example:"bastille0" format:"string"`
	Gtwip      string `json:"gtwip,omitempty" example:"n.n.n.n" format:"string" validate:"omitempty,ip|cidr"`
	Ipip       string `json:"ipip,omitempty" example:"n.n.n.n,i.i.i.i" format:"string"`
	Value      string `json:"value,omitempty" example:"99" format:"string" validate:"omitempty,number"`
	Vlanid     string `json:"vlanid,omitempty" example:"vlan10" format:"string"`
	Zfsoptions string `json:"zfsoptions,omitempty" example:"" format:"string"`
}

type destroyModel struct {
	Options     string `json:"options,omitempty" example:"-a|-c|-f|-y|-x" format:"string"`
	JailRelease string `json:"jail|release" example:"jail_target or release" format:"string" validate:"required"`
}

type editModel struct {
	Options string `json:"options,omitempty" example:"-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
	File    string `json:"file,omitempty" example:"file_name" format:"string"`
}

type etcupdateModel struct {
	Options         string `json:"options,omitempty" example:"-d|-f|-x" format:"string"`
	Bootstraptarget string `json:"bootstrap|target" example:"jail_name|bootstrap" format:"string" validate:"required"`
	Action          string `json:"action,omitempty" example:"diff|resolve|update" format:"string"`
	Release         string `json:"release,omitempty" example:"14.3-RELEASE" format:"string"`
}

type exportModel struct {
	Options string `json:"options,omitempty" example:"-a|--gz|-r|-s|--tgz|--txz|-v|--xz|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Path    string `json:"path" example:"/host/path" format:"string" validate:"filepath"`
}

type htopModel struct {
	Options string `json:"options,omitempty" example:"-a|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
}

type importModel struct {
	Options string `json:"options,omitempty" example:"-f|-M|-v|-x" format:"string"`
	File    string `json:"file" example:"/path/to/archive.file" format:"string" validate:"filepath"`
	Release string `json:"release,omitempty" example:"release_name" format:"string"`
}

type jcpModel struct {
	Options    string `json:"options,omitempty" example:"-q|-x" format:"string"`
	Sourcejail string `json:"source_jail" example:"source_jail" format:"string" validate:"required"`
	Jailpath   string `json:"jail_path" example:"/source_jail/path" format:"string" validate:"filepath"`
	Destjail   string `json:"dest_jail" example:"dest_jail" format:"string" validate:"required"`
	Jailpath2  string `json:"jail_path2" example:"/dest_jail/path" format:"string" validate:"filepath"`
}

type limitsModel struct {
	Options string `json:"options,omitempty" example:"-a|-l|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Action  string `json:"action" example:"add|remove|clear|reset|list|show|active|stats" format:"string" validate:"required"`
	Option  string `json:"option,omitempty" example:"see rctl" format:"string"`
	Value   string `json:"value,omitempty" example:"depends on Option" format:"string"`
}

type listModel struct {
	Options string `json:"options,omitempty" example:"-d|-j|-p|-u|-x" format:"string"`
	Action  string `json:"action" example:"all|backup|export|import|ip|jail|limit|log|path|port|prio|release|state|template|type" format:"string" validate:"required"`
}

type migrateModel struct {
	Options string `json:"options,omitempty" example:"-a|-b|-d|--doas|-l|-p|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Remote  string `json:"remote" example:"user@host:port" format:"string" validate:"required"`
}

type monitorModel struct {
	Options string `json:"options,omitempty" example:"-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Action  string `json:"action" example:"enable|disable|status|add|delete|list" format:"string" validate:"required"`
	Service string `json:"service,omitempty" example:"nginx" format:"string"`
}

type mountModel struct {
	Options        string `json:"options,omitempty" example:"-a|-x" format:"string"`
	Target         string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Hostpath       string `json:"host_path" example:"/host/path" format:"string" validate:"filepath"`
	Jailpath       string `json:"jail_path" example:"/jail/path" format:"string" validate:"filepath"`
	Filesystemtype string `json:"filesystem_type" example:"tmpfs|nullfs" format:"string" validate:"required"`
	Option         string `json:"option" example:"ro|rw|rw,nosuid,mode=01777" format:"string" validate:"required"`
	Dump           string `json:"dump" example:"0" format:"string" validate:"required"`
	Passnumber     string `json:"pass_number" example:"0" format:"string" validate:"required"`
}

type networkModel struct {
	Options string `json:"options,omitempty" example:"-a|-B|-M|-n|-P|-V|-v|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Action  string `json:"action" example:"add|remove" format:"string" validate:"required"`
	Iface   string `json:"interface,omitempty" example:"bastille0" format:"string"`
	Ip      string `json:"ip,omitempty" example:"n.n.n.n" format:"string" validate:"omitempty,ip"`
	Vlanid  string `json:"vlanid,omitempty" example:"vlan10" format:"string"`
}

type pkgModel struct {
	Options string `json:"options,omitempty" example:"-a|-H|-y|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Arg     string `json:"args" example:"install pkg_name" format:"string" validate:"required"`
}

type rcpModel struct {
	Options  string `json:"options,omitempty" example:"-q|-x" format:"string"`
	Target   string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Jailpath string `json:"jail_path" example:"/jail/path" format:"string" validate:"filepath"`
	Hostpath string `json:"host_path" example:"/host/path" format:"string" validate:"filepath"`
}

type rdrModel struct {
	Options      string `json:"options,omitempty" example:"-d|-i|-s|-t|-x" format:"string"`
	Odestination string `json:"odestination,omitempty" example:"depends on options destination" format:"string"`
	Ointerface   string `json:"ointerface,omitempty" example:"depends on options interface" format:"string"`
	Osource      string `json:"osource,omitempty" example:"depends on options source" format:"string"`
	Otype        string `json:"otype,omitempty" example:"depends on options type" format:"string"`
	Target       string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Action       string `json:"action" example:"clear|reset|list|tcp|udp" format:"string" validate:"required"`
	Hostport     string `json:"host_port,omitempty" example:"3000" format:"string" validate:"omitempty,port"`
	Jailport     string `json:"jail_port,omitempty" example:"3000" format:"string" validate:"omitempty,port"`
	Log          string `json:"log,omitempty" example:"log" format:"string"`
	Logopts      string `json:"logopts,omitempty" example:"log options" format:"string"`
}

type renameModel struct {
	Options string `json:"options,omitempty" example:"-a|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Newname string `json:"new_name" example:"new_name" format:"string" validate:"required"`
}

type restartModel struct {
	Options string `json:"options,omitempty" example:"-b|-d|-v|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Value   string `json:"value,omitempty" example:"seconds to -d option" format:"string" validate:"omitempty,number"`
}

type serviceModel struct {
	Options     string `json:"options,omitempty" example:"-a|-x" format:"string"`
	Target      string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Servicename string `json:"service_name" example:"nginx" format:"string" validate:"required"`
	Args        string `json:"args" example:"start|restart|stop|enable|disable" format:"string" validate:"required"`
}

type setupModel struct {
	Options string `json:"options,omitempty" example:"-y|-x" format:"string"`
	Action  string `json:"action" example:"bridge|loopback|pf|firewall|shared|vnet|storage" format:"string" validate:"required"`
}

type startModel struct {
	Options string `json:"options,omitempty" example:"-b|-d|-v|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Value   string `json:"value,omitempty" example:"seconds to -d option" format:"string" validate:"omitempty,number"`
}

type stopModel struct {
	Options string `json:"options,omitempty" example:"-v|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
}

type sysrcModel struct {
	Options string `json:"options,omitempty" example:"-a|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Args    string `json:"args" example:"sshd_enable=\"YES\"" format:"string" validate:"required"`
}

type tagsModel struct {
	Options string `json:"options,omitempty" example:"-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Action  string `json:"action" example:"add|delete|list" format:"string" validate:"required"`
	Tgs     string `json:"tags,omitempty" example:"tag_name1,tag_name2" format:"string"`
}

type templateModel struct {
	Options  string `json:"options,omitempty" example:"-a|-x" format:"string"`
	Target   string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Action   string `json:"action,omitempty" example:"--convert" format:"string"`
	Template string `json:"template" example:"template/name" format:"string" validate:"required"`
}

type topModel struct {
	Options string `json:"options,omitempty" example:"-a|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
}

type umountModel struct {
	Options  string `json:"options,omitempty" example:"-a|-x" format:"string"`
	Target   string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Jailpath string `json:"jail_path" example:"/jail/path" format:"string" validate:"required"`
}

type updateModel struct {
	Options string `json:"options,omitempty" example:"-a|-f|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
}

type upgradeModel struct {
	Options string `json:"options,omitempty" example:"-a|-f|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Action  string `json:"NEW_RELEASE|install" example:"n.n-RELEASE|install" format:"string" validate:"required"`
}

type verifyModel struct {
	Options string `json:"options,omitempty" example:"-x" format:"string"`
	Action  string `json:"RELEASE|TEMPLATE" example:"n.n-RELEASE|template_name" format:"string" validate:"required"`
}

type zfsModel struct {
	Options     string `json:"options,omitempty" example:"-a|-v|-x" format:"string"`
	Target      string `json:"target" example:"jail_target" format:"string" validate:"required"`
	Action      string `json:"action" example:"destroy|rollback|snapshot|df|usage|get|set|jail|unjail" format:"string" validate:"required"`
	Tag         string `json:"tag,omitempty" example:"tag_name" format:"string"`
	Key         string `json:"key,omitempty" example:"name for set action" format:"string"`
	Value       string `json:"value,omitempty" example:"value for set action" format:"string"`
	Pooldataset string `json:"pool/dataset,omitempty" example:"pool/dataset" format:"string"`
	Jailpath    string `json:"/jail/path,omitempty" example:"/jail/path" format:"string"`
}
