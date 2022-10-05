package middleware

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/lrth06/go-chat/lib/utils/config"
)

// ISSUE #13 This should be refactored such that admin is its own middleware, then self check happens in the handler
func SelfOrAdmin(c *fiber.Ctx) error {
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
			fmt.Println("user does not have admin or moderator role")
			if claims["id"] != c.Params("id") {
				return c.Status(401).SendString("Unauthorized")
			}
		}
	}
	return c.Next()
}
