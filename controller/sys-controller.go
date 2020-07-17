package controller

import (
	"strconv"

	. "github.com/cacing69/api/lib"
	"github.com/gofiber/fiber"
)

func SysIndex(c *fiber.Ctx) {
	// panic("normally this would crash your app")
	c.JSON(Res{
		Message: "welcome to api-source",
		Status:  true,
		Code:    200,
	})
}

func SysPing(c *fiber.Ctx) {
	c.JSON(Res{
		Message: "pong!",
		Status:  true,
		Code:    200,
	})
}

func SysError(c *fiber.Ctx) {

	msg := []string{
		"lorem ipsum",
		"dolor sit amet",
	}

	code := c.Params("code")
	_code, err := strconv.Atoi(code)
	if err != nil {
		c.Next(err)
		return
	}

	c.JSON(Res{
		Message: "Error Testing",
		Code:    _code,
		Errors:  msg,
	})
}

func SysValidate(c *fiber.Ctx) {
	type req struct {
		Name  string `form:"name" json:"name" validate:"required"`
		Email string `json:"email"  json:"email" validate:"required,email"`
	}

	r := new(req)

	err, validate := Validate(c, r)

	if err != nil {
		ValidateFailed(c, err, validate)
		return
	}

	c.JSON(Res{
		Data:    r,
		Message: "success",
		Status:  true,
		Code:    200,
	})
}

// func SysQuery(c *fiber.Ctx) {
// 	_id, _err := strconv.Atoi(c.Query("id"))

// 	if _err != nil {
// 		c.Next(_err)
// 		return
// 	}

// 	_row := dbq.MustQ(c.Context(), conf.DB, "SELECT * FROM m_user where user_id = ?", entity.UserSingleOption(), _id)

// 	if _row == nil {
// 		NotFound(c)
// 		return
// 	} else {
// 		c.JSON(Res{
// 			Data:    _row,
// 			Message: "success",
// 			Status:  true,
// 			Code:    200,
// 		})
// 	}
// }
