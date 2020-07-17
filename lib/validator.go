package lib

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
)

func Validate(c *fiber.Ctx, r interface{}) (error, []string) {

	validate := validator.New()

	c.BodyParser(r)
	err := validate.Struct(r)

	if err != nil {

		trace := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			var field string = strings.ToLower(e.Field())
			switch e.Tag() {
			case "required":
				trace = append(trace, fmt.Sprintf("%s is required", field))
			case "email":
				trace = append(trace, fmt.Sprintf("%s is not valid email", field))
			case "gte":
				trace = append(trace, fmt.Sprintf("%s value must be greater than %s", field, e.Param()))
			case "lte":
				trace = append(trace, fmt.Sprintf("%s value must be lower than %s", field, e.Param()))
			}
		}
		return errors.New("validation failed"), trace
	} else {
		return nil, nil
	}

}
