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
	body := new(structs.User)
	c.BodyParser(&body)

	err := Validator.Struct(body)
	// [ ] need more concise errors
	errMsg :=""
	if err != nil {
		errs := err.(validator.ValidationErrors)
		errMsg = fmt.Sprintf("Unble to validate %s", errs[0].Field())
		return c.Status(422).JSON(fiber.Map{
			"msg": errMsg,
		})
	}
	return c.Next()
}
