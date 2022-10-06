package users

import (
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

	env, err := config.GetConfig()
	if err != nil {
		return c.SendStatus(400)
	}
	secret := env.TokenSecret
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.SendStatus(400)
	}
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(400)
	}
	//find and update in db, then return new jwt token with updated user information
	query := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: user}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	if err := config.ConnDB("Users").FindOneAndUpdate(c.Context(), query, update, opts).Decode(&user); err != nil {
		return c.SendStatus(400)
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
		return c.SendStatus(400)
	}
	return c.JSON(fiber.Map{
		"msg":   "User updated successfully!",
		"token": jwt,
	})
}
