package validation

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ValidateUser(c *fiber.Ctx) error {
	fmt.Println("Validating user...")

	// FIXME: Currently Returns NIL, SHOULD BE IMPLEMENTED ASAP
	return c.Next()
}
