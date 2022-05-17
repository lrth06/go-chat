package users

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/models"
	"github.com/lrth06/go-chat/lib/utils/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	userId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "Invalid ID."})
	}
	query := bson.D{{Key: "_id", Value: userId}}
	user := models.User{}
	if err := config.ConnDB("Users").FindOne(c.Context(), query).Decode(&user); err != nil {
		return c.Status(404).JSON(fiber.Map{"msg": "User not found."})
	}
	if _, err := config.ConnDB("Users").DeleteOne(c.Context(), query); err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "User not deleted."})
	}

	//TODO: add production delete
	//delete users images directory
	if err := os.RemoveAll("./tmp/uploads/users/" + idParam); err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "User images not deleted."})
	}
	return c.JSON(fiber.Map{
		"msg": "User deleted successfully!",
	})
}
