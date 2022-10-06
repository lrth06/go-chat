package middleware

import (
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AdminCheck(c *fiber.Ctx) error {
	// use c.Locals(user) to get the user from the middleware
	// if the user is not an admin, set c.Locals("admin", false)
	// if the user is an admin, set c.Locals("admin", true)
	// then in the handler, check if c.Locals("admin") is true or false
	if c.Locals("userclaims") == (nil) {
		c.Locals("admin", false)
		return c.Next()
	}
	claims := c.Locals("userclaims").(jwt.MapClaims)
	role := base64.StdEncoding.EncodeToString([]byte("site:admin"))
	role2 := base64.StdEncoding.EncodeToString([]byte("site:moderator"))
	for _, r := range claims["roles"].([]interface{}) {
		if r == role || r == role2 {
			c.Locals("admin", true)
			return c.Next()
		}
	}
	c.Locals("admin", false)
	return c.Next()
}
