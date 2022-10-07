package rooms

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/models"
	"github.com/lrth06/go-chat/lib/utils/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateRoom(c *fiber.Ctx) error {
	room := models.Room{}
	room.ID = primitive.NewObjectID()
	if err := c.BodyParser(&room); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": "Invalid request.",
		})
	}
	room.Owner, _ = primitive.ObjectIDFromHex(c.Locals("id").(string))

	// [ ] VALIDATE ROOM DATA HERE
	//check for duplicate room name
	filter := bson.M{
		"name": room.Name,
	}
	if err := config.ConnDB("Rooms").FindOne(c.Context(), filter).Decode(&room); err == nil {
		return c.Status(409).JSON(fiber.Map{
			"msg": "Room already exists!",
		})
	}

	if _, err := config.ConnDB("Rooms").InsertOne(c.Context(), room); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"msg": "Failed to create room!",
		})
	}
	return c.JSON(room)

}
