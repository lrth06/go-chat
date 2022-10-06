package users

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/lrth06/go-chat/lib/models"
	"github.com/lrth06/go-chat/lib/utils/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateUser(c *fiber.Ctx) error {

	fmt.Println("Self:",c.Locals("self"),"Admin:", c.Locals("admin"))
	if c.Locals("admin") == false && c.Locals("self") == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg": "Unauthorized",
		})
	}

	env, err := config.GetConfig()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": "Server error."})
	}
	secret := env.TokenSecret
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": "Server error."})
	}
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": "Server error."})
	}
	//find and update in db, then return new jwt token with updated user information
	query := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: user}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	if err := config.ConnDB("Users").FindOneAndUpdate(c.Context(), query, update, opts).Decode(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": "Server error."})
	}
	jwtClaims := jwt.MapClaims{
		"id":     user.ID,
		"name":   user.Name,
		"email":  user.Email,
		"avatar": user.Avatar,
		"roles":  user.Roles,
		"exp":    jwt.TimeFunc().Add(time.Hour * 72).Unix(),
	}
	jwt, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims).SignedString([]byte(secret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": "Server error."})
	}
	return c.JSON(fiber.Map{
		"msg":   "User updated successfully!",
		"token": jwt,
	})
}
