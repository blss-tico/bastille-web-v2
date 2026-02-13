package config

var BastilleM BastilleModel
var AddrModel string
var BwAddrModel string
var BwPortModel string
var BwUsers []UsersModel
var KeyModel string
var JwtKeyModel []byte
var RefreshKeyModel []byte
var NodesListModel []NodesModel

type UsersModel struct {
	ID       string `json:"id,omitempty" example:"1" format:"string"`
	Username string `json:"username" example:"user" format:"string"`
	Password string `json:"password" example:"secretpassword" format:"string"`
}

type NodesModel struct {
	Nodename string `json:"nodename" example:"Node01" format:"string"`
	Nodeip   string `json:"nodeip" example:"192.168.1.1" format:"string"`
	Nodeport string `json:"nodeport" example:"8007" format:"string"`
}

type BastilleOptionsModel struct {
	Sflag string
	Lflag string
	Text  string
}

type BastilleCommandsModel struct {
	Command     string                 `json:"command"`
	Description string                 `json:"description"`
	Options     []BastilleOptionsModel `json:"options"`
	Fields      []string               `json:"fields"`
	Help        string                 `json:"help"`
	HelpUrl     string                 `json:"helpUrl"`
}

type BastilleModel struct {
	Software string                  `json:"software"`
	Options  []string                `json:"options"`
	Help     string                  `json:"help"`
	Commands []BastilleCommandsModel `json:"commands"`
}
