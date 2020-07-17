package controller

import (
	"time"

	"github.com/cacing69/api/entity"
	"github.com/cacing69/api/lib"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

var LOGIN_EXPIRATION_DURATION = time.Duration(5) * time.Minute
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("secret")
var APPLICATION_NAME = "api-source"

func AuthToken(c *fiber.Ctx) {
	authorization := c.Fasthttp.Request.Header.Peek("Authorization")

	username, password, state := ParseBasicAuth(string(authorization))

	if state {
		auth, user := authenticate(username, password)
		if auth {

			signedToken, err := claimToken(user)

			if err != nil {
				c.Next(err)
				return
			}

			c.JSON(Res{
				Message: "auth token success",
				Data: M{
					"token": signedToken,
				},
			})
		} else {
			lib.BadRequest(c, "invalid username / password")
		}
	} else {
		lib.BadRequest(c, "invalid authorization")
	}
}

func claimToken(user entity.User) (string, error) {
	claims := JwtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Id: user.Id,
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	return token.SignedString(JWT_SIGNATURE_KEY)
}

func authenticate(username, password string) (bool, entity.User) {
	if username == "cacing69" && password == "23Cacing09#@" {
		return true, entity.User{
			Id:   1,
			Name: "cacing69",
		}
	} else {
		return false, entity.User{}
	}
}
