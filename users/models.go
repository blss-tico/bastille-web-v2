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
