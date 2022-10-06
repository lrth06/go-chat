package rooms

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/models"
	"github.com/lrth06/go-chat/lib/utils/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetRooms(c *fiber.Ctx) error {
	query := bson.D{{}}
	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := config.ConnDB("Rooms").Find(c.Context(), query, opts)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var rooms []models.Room = make([]models.Room, 0)
	if err := cursor.All(c.Context(), &rooms); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(rooms)
}
