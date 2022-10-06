package permissions

import (
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {

	return c.Next()
}

func RequireAuth(c *fiber.Ctx) error {
	reject := func(c *fiber.Ctx) error {
		// c.SendFile("client/build/index.html")
		c.Status(401).Send([]byte("Unauthorized"))

		return nil
	}
	if c.Get("Authorization") == "" {
		return reject(c)
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
