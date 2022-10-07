package rooms

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/models"
	"github.com/lrth06/go-chat/lib/utils/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateRoom(c *fiber.Ctx) error {

	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": "Invalid room ID",
		})
	}
	room := models.Room{}
	if err := c.BodyParser(&room); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": "Invalid room data",
		})
	}

	//check if c.Locals("userclaims") id is the owner of the room or in settings.moderators

	// [ ] check if room exists
	// [ ] check if user is owner or moderator of room

	query := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: room}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	if err := config.ConnDB("Rooms").FindOneAndUpdate(c.Context(), query, update, opts).Decode(&room); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"msg": "Room not found",
		})
	}

	return c.JSON(fiber.Map{
		"msg":  "Room updated successfully!",
		"room": room,
	})
}
