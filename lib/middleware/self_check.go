package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func SelfCheck(c *fiber.Ctx) error {
	claims := c.Locals("userclaims").(jwt.MapClaims)
	if claims["id"] == c.Params("id") {
		c.Locals("self", true)
		return c.Next()
	}
	c.Locals("self", false)
	return c.Next()
}
