package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/models"
	"github.com/lrth06/go-chat/lib/utils/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	userId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.SendStatus(400)
	}
	query := bson.D{{Key: "_id", Value: userId}}
	opts := options.FindOne().SetProjection(bson.D{{Key: "password", Value: 0}, {Key: "token", Value: 0}})
	user := models.User{}
	if err := config.ConnDB("Users").FindOne(c.Context(), query, opts).Decode(&user); err != nil {
		return c.Status(404).SendString("User not found!")
	}
	return c.JSON(user)
}
