package web

import (
	"bastille-web-v2/config"

	"github.com/dgrijalva/jwt-go"
)

type userModel struct {
	ID       int    `json:"id,omitempty" example:"1" format:"string"`
	Username string `json:"username" example:"user" format:"string"`
	Password string `json:"password" example:"secretpassword" format:"string"`
}

type claimsModel struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type templatesModel struct {
	CommandName string
	Data        config.BastilleModel
}
