package main

import (
	"github.com/cacing69/api/conf"
	"github.com/cacing69/api/controller"
	. "github.com/cacing69/api/lib"
	"github.com/gofiber/fiber"
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
				c.Status(code).JSON(
					Res{
						Message: err.Error(),
					},
				)
			},
		},
	)

	// default router
	r.Get("/", controller.SysIndex)

	// system standard for testing route
	sys := r.Group("/sys")
	sys.Get("/ping", controller.SysPing)
	sys.Get("/error/:code", controller.SysError)
	sys.Get("/query", controller.SysQuery)
	sys.Post("/validate", controller.SysValidate)

	r.Listen(3000)
}
