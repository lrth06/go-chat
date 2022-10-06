package rooms

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/models"
	"github.com/lrth06/go-chat/lib/utils/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetRoom(c *fiber.Ctx) error {
	idParam := c.Params("id")
	roomId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": "Invalid room ID",
		})
	}
	opts := options.FindOne().SetProjection(bson.D{{Key: "password", Value: 0}})
	query := bson.D{{Key: "_id", Value: roomId}}
	room := models.Room{}
	if err := config.ConnDB("Rooms").FindOne(c.Context(), query, opts).Decode(&room); err != nil {
		fmt.Println(err)
		return c.Status(404).JSON(fiber.Map{"msg": "Room not found."})
	}
	return c.JSON(room)
}
