package lib

import "github.com/gofiber/fiber"

type Res struct {
	Status  bool        `json:"status" default:"true"`
	Message string      `json:"message" default:"success"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
	Errors  []string    `json:"errors"`
	Code    int         `json:"code" default:"200"`
}

func (r Res) JSON(c *fiber.Ctx) error {
	return c.Status(r.Code).JSON(r)
}

func NotFound(c *fiber.Ctx) error {
	return Res{
		Message: "not found",
		Code:    404,
	}.JSON(c)
}

func ValidateFailed(c *fiber.Ctx, err error, validate []string) error {
	return Res{
		Message: err.Error(),
		Code:    422,
		Errors:  validate,
	}.JSON(c)
}

func BadRequest(c *fiber.Ctx, message string) error {
	return Res{
		Message: message,
		Code:    422,
	}.JSON(c)
}
