package web

import (
	"bastille-web-v2/config"

	"github.com/dgrijalva/jwt-go"
)

type claimsModel struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type templatesModel struct {
	CommandName string
	Data        config.BastilleModel
}
