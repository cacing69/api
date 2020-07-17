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

func ResNotFound(c *fiber.Ctx) error {
	return c.Status(404).JSON(
		Res{
			Message: "data not found",
			Code:    404,
		},
	)
}

func ResValidate(c *fiber.Ctx, err error, validate []string) error {
	return c.Status(422).JSON(
		Res{
			Message: err.Error(),
			Code:    422,
			Errors:  validate,
		},
	)
}

func ResErr(c *fiber.Ctx, res Res) error {
	return c.Status(res.Code).JSON(
		res,
	)
}
