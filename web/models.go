package web

import (
	"bastille-web-v2/config"
)

type templatesModel struct {
	CommandName string
	Data        config.BastilleModel
}
