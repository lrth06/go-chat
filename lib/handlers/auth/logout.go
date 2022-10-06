package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/lrth06/go-chat/lib/utils/config"
)

func LogoutUser(c *fiber.Ctx) error {
	env, err := config.GetConfig()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": "Server error."})
	}
	secret := env.TokenSecret
	jwt, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": jwt.TimeFunc().Add(-5 * time.Minute).Unix(),
	}).SignedString([]byte(secret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": "Server error."})
	}
	return c.JSON(fiber.Map{
		"msg":   "User logged out successfully!",
		"token": jwt,
	})
}
