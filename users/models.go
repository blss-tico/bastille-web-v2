package users

import "github.com/dgrijalva/jwt-go"

type claimsModel struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type AllUsersModel struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type loginLogoutModel struct {
	Bw_actk    string `json:"bw_actk"`
	Bw_rftk    string `json:"bw_rftk"`
	Request_id string `json:"request_id"`
}
