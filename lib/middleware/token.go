package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/lrth06/go-chat/lib/utils/config"
)

func ExtractToken(c *fiber.Ctx) error {
	env, err := config.GetConfig()
	if err != nil {
		return c.SendStatus(400)
	}
	secret := env.TokenSecret
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(401).SendString("Unauthorized")
	}
	token = strings.Split(token, " ")[1]
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return c.Status(401).JSON(map[string]string{
			"message": "Invalid token",
		})
	}
	c.Set("id", claims["id"].(string))
	c.Set("token", token)
	return c.Next()
}
