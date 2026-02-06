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
	Id       int
	Username string
	Password string
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

type NodesModel struct {
	Host string `json:"host"`
	Ip   string `json:"ip"`
	Port string `json:"port"`
	Key  string `json:"key"`
}
