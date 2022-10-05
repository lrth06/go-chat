package middleware

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/lrth06/go-chat/lib/models"
	"github.com/lrth06/go-chat/lib/utils/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SelfAdminorMod(c *fiber.Ctx) error {
	env, err := config.GetConfig()
	if err != nil {
		return c.SendStatus(400)
	}
	secret := env.TokenSecret
	auth := c.Get("Authorization")
	if auth == "" {
		return c.Status(401).SendString("Unauthorized")
	}
	token := strings.Split(auth, " ")[1]
	if token == "" {
		return c.Status(401).SendString("Unauthorized")
	}
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return c.Status(401).SendString("Unauthorized")
	}

	role := base64.StdEncoding.EncodeToString([]byte("site:admin"))
	role2 := base64.StdEncoding.EncodeToString([]byte("site:moderator"))
	for _, r := range claims["roles"].([]interface{}) {
		if r != role || r != role2 {
			fmt.Println("user does not have site admin or moderator role")
			idParam := c.Params("id")
			roomId, err := primitive.ObjectIDFromHex(idParam)
			if err != nil {
				return c.SendStatus(400)
			}
			opts := options.FindOne().SetProjection(bson.D{{Key: "password", Value: 0}})
			query := bson.D{{Key: "_id", Value: roomId}}
			room := models.Room{}
			if err := config.ConnDB("Rooms").FindOne(c.Context(), query, opts).Decode(&room); err != nil {
				fmt.Println(err)
				return c.Status(404).SendString("room not found!")
			}
			id := claims["id"].(string)
			ownerId := room.Owner.Hex()
			if id != ownerId {
				for _, m := range room.Settings.Moderators {
					if m.Hex() == id {
						fmt.Println("user is room moderator")
						return c.Next()
					}
				}
			}
		}
	}
	fmt.Println("user is room owner or site admin")
	return c.Next()

}
