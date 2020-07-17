package main

import (
	"strings"

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
				code := 400

				c.JSON(Res{
					Message: strings.ToLower(err.Error()),
					Code:    code,
				})
			},
		},
	)

	// Match all routes starting with /api
	r.Use(JwtMiddleware)

	r.Use(middleware.Recover())
	r.Use(middleware.Logger())

	// default router
	r.Get("/", controller.SysIndex)

	// auth endpoint
	auth := r.Group("/auth")
	auth.Post("/token", controller.AuthToken)
	auth.Get("/check", controller.AuthCheck)

	// system standard for testing route
	sys := r.Group("/sys")
	sys.Get("/ping", controller.SysPing)
	sys.Get("/error/:code", controller.SysError)
	sys.Post("/validate", controller.SysValidate)
	// sys.Get("/query", controller.SysQuery)

	// user endpoint
	user := r.Group("/user")
	// user.Get("/orm", controller.UserIndexOrm)
	user.Get("", controller.UserIndex)
	// user.Get("/:id", controller.UserShow)
	user.Post("", controller.UserStore)
	// user.Patch("/:id", controller.UserUpdate)
	// user.Delete("", controller.UserDelete)

	r.Use(func(c *fiber.Ctx) {
		NotFound(c)
	})

	r.Listen(3000)
}
