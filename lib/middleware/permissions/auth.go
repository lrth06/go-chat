package permissions

import (
	"github.com/gofiber/fiber/v2"
)

func RequireAuth(c *fiber.Ctx) error {
	if c.Get("Authorization") == "" {
		return c.Status(401).JSON(fiber.Map{
			"msg": "Unauthorized",
		})
	}
	return c.Next()
}

func SocketAuth(c *fiber.Ctx) error {
	token := c.Query("token")
	if token == "" {
		c.Status(401).Send([]byte("Unauthorized"))
		return nil
	}

	return c.Next()
}
