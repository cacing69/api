package controller

import (
	"time"

	"github.com/cacing69/api/entity"
	. "github.com/cacing69/api/lib"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("secret")

func AuthToken(c *fiber.Ctx) {
	authorization := c.Fasthttp.Request.Header.Peek("Authorization")

	username, password, state := ParseBasicAuth(string(authorization))

	if state {
		auth, _ := authenticate(username, password)
		if auth {
			Res{
				Message: "auth token success",
				Data: fiber.Map{
					"username": username,
					"password": password,
				},
			}.JSON(c)
		} else {
			Res{
				Message: "invalid username / password",
			}.JSON(c)
		}
	} else {
		Res{
			Message: "invalid authorization",
			Code:    400,
		}.JSON(c)
	}
}

func authenticate(username, password string) (bool, entity.User) {

	// fake auth
	if username == "cacing69" && password == "23Cacing09#@" {
		return true, entity.User{
			Id:   1,
			Name: "cacing69",
		}
	} else {
		return false, entity.User{}
	}
}
