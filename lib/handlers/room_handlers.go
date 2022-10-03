package handlers

import "github.com/gofiber/fiber/v2"

func CreateRoom(c *fiber.Ctx) error {
	return c.Status(500).JSON(map[string]string{
		"msg": "Not implemented",
	})
}

func GetRoom(c *fiber.Ctx) error {
	return c.Status(500).JSON(map[string]string{
		"msg": "Not implemented",
	})
}

func UpdateRoom(c *fiber.Ctx) error {
	return c.Status(500).JSON(map[string]string{
		"msg": "Not implemented",
	})
}

func DeleteRoom(c *fiber.Ctx) error {
	return c.Status(500).JSON(map[string]string{
		"msg": "Not implemented",
	})
}

func GetRooms(c *fiber.Ctx) error {
	return c.Status(500).JSON(map[string]string{
		"msg": "Not implemented",
	})
}
