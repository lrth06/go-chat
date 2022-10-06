package token

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/lrth06/go-chat/lib/utils/config"
)

func ExtractToken(c *fiber.Ctx) error {
	fmt.Println("Extracting Token")
	env, err := config.GetConfig()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": "Server error."})
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
			"msg": "Invalid token",
		})
	}
	userclaims := jwt.MapClaims{
		"id":    claims["id"],
		"roles": claims["roles"],
		"exp":   claims["exp"],
	}
	c.Locals("id", claims["id"])
	c.Locals("userclaims", userclaims)

	return c.Next()
}
