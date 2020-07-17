package lib

import (
	"github.com/dgrijalva/jwt-go"
)

type M map[string]interface{}

type JwtClaims struct {
	jwt.StandardClaims
	Id int `json:"id"`
}
