package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/structs"
	"github.com/lrth06/go-chat/lib/utils"
)

func Logger(c *fiber.Ctx) error {
	payload := structs.Payload{
		Status:  c.Response().StatusCode(),
		Method:  c.Method(),
		Path:    c.Path(),
		IP:      c.IP(),
		Headers: c.GetReqHeaders(),
	}

	utils.LogItem("INFO", payload)
	return c.Next()
}
