package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/lrth06/go-chat/lib/structs"

	"github.com/gofiber/fiber/v2"
)

type IError struct {
	Field string
	Tag   string
	Value string
}

var Validator = validator.New()

func ValidateUser(c *fiber.Ctx) error {
	fmt.Println("Validating user...")
	var errors []*IError
	body := new(structs.User)
	c.BodyParser(&body)

	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return c.Next()
}
