package permissions

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/lrth06/go-chat/lib/models"
	"github.com/lrth06/go-chat/lib/utils/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func OwnerModerator(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": "Invalid room ID",
		})
	}
	// allow if either is true (owner or moderator)
	claims := c.Locals("userclaims").(jwt.MapClaims)
	userId, err := primitive.ObjectIDFromHex(claims["id"].(string))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": "Invalid user ID",
		})
	}

	existingRoom := models.Room{}
	err = config.ConnDB("Rooms").FindOne(c.Context(), bson.M{"_id": id}).Decode(&existingRoom)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": "Room not found",
		})
	}
	fmt.Println(
		existingRoom.Owner,
		userId,
	)

	authorized := false
	if userId == existingRoom.Owner {
		fmt.Println("User is Room Owner")
		authorized = true
	}
	for _, m := range existingRoom.Settings.Moderators {
		if m == userId {
			fmt.Println("User is moderator")
			authorized = true
		}
	}
	if !authorized {

		return c.Status(401).JSON(fiber.Map{
			"msg": "Unauthorized",
		})
	}

	return c.Next()
}
