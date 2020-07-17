package lib

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

var LOGIN_EXPIRATION_DURATION = time.Duration(24) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("secret")
var APPLICATION_NAME = "api-source"

func JwtMiddleware(c *fiber.Ctx) {
	if IsPublicEndpoint(c.Path()) {
		c.Next()
	} else {
		if c.Path() == "/" {
			c.Next()
		} else {
			bearer := c.Fasthttp.Request.Header.Peek("Authorization")

			if len(bearer) == 0 {
				UnAuthorized(c, []string{"empty token"})
				return
			}

			if !strings.Contains(string(bearer), "Bearer") {
				UnAuthorized(c, []string{"invalid token"})
				return
			}

			authorization := strings.Replace(string(bearer), "Bearer ", "", -1)

			token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {

				if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("signing method invalid")

				} else if method != JWT_SIGNING_METHOD {
					return nil, fmt.Errorf("signing method invalid")
				}
				return JWT_SIGNATURE_KEY, nil
			})

			if err != nil {
				c.Next(err)
				return
			}

			claims, state := token.Claims.(jwt.MapClaims)

			if !state || !token.Valid {
				c.Next(err)
				return
			}

			c.Fasthttp.SetUserValue("auth", claims)

			c.Next()
		}
	}
}
