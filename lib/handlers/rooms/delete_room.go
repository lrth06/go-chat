package rooms

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/models"
	"github.com/lrth06/go-chat/lib/utils/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteRoom(c *fiber.Ctx) error {
	idParam := c.Params("id")
	roomId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "Invalid ID."})
	}
	query := bson.D{{Key: "_id", Value: roomId}}
	room := models.Room{}
	if err := config.ConnDB("Rooms").FindOne(c.Context(), query).Decode(&room); err != nil {
		return c.Status(404).JSON(fiber.Map{"msg": "Room not found."})
	}
	if _, err := config.ConnDB("Rooms").DeleteOne(c.Context(), query); err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": "Room not deleted."})
	}
	return c.JSON(fiber.Map{
		"msg": "Room deleted successfully!",
	})
}
