package controller

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber"
)

func UserIndex(c *fiber.Ctx) {
	spew.Dump(c.Fasthttp.UserValue("auth"))
}

func UserShow(c *fiber.Ctx) {

}

func UserStore(c *fiber.Ctx) {

}

func UserUpdate(c *fiber.Ctx) {

}

func UserDelete(c *fiber.Ctx) {

}
