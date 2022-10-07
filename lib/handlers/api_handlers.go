package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetRandomID(c *fiber.Ctx) error {
	id := uuid.New().String()
	return c.JSON(fiber.Map{
		"id": id,
	})
}
