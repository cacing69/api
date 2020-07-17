package main

import (
	"fmt"

	"github.com/cacing69/api/conf"
	"github.com/cacing69/api/controller"
	. "github.com/cacing69/api/lib"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main() {
	defer conf.DB.Close()

	r := fiber.New(
		&fiber.Settings{
			Prefork:       true,
			CaseSensitive: true,
			StrictRouting: true,
			ErrorHandler: func(c *fiber.Ctx, err error) {
				code := fiber.StatusInternalServerError

				Res{
					Message: err.Error(),
					Code:    code,
				}.JSON(c)
			},
		},
	)

	// Match all routes starting with /api
	r.Use(func(c *fiber.Ctx) {
		if IsPublicEndpoint(c.Path()) {
			c.Next()
			fmt.Println("🥇 public endpoint")
		} else {
			fmt.Println("🥇 private endpoint secure with jwt")
		}
	})

	r.Use(middleware.Recover())

	// default router
	r.Get("/", controller.SysIndex)

	// auth endpoint
	auth := r.Group("/auth")
	auth.Post("/token", controller.AuthToken)

	// system standard for testing route
	sys := r.Group("/sys")
	sys.Get("/ping", controller.SysPing)
	sys.Get("/error/:code", controller.SysError)
	sys.Post("/validate", controller.SysValidate)
	sys.Get("/query", controller.SysQuery)

	// user endpoint
	user := r.Group("/user")
	user.Get("", controller.UserIndex)
	user.Get("/:id", controller.UserShow)
	user.Post("", controller.UserStore)
	user.Patch("/:id", controller.UserUpdate)
	user.Delete("", controller.UserDelete)

	r.Use(func(c *fiber.Ctx) {
		NotFound(c)
	})

	r.Listen(3000)
}
