package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/lrth06/go-chat/lib/models"
	"github.com/lrth06/go-chat/lib/utils/config"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *fiber.Ctx) error {
	// Get user from request body
	env, err := config.GetConfig()
	if err != nil {
		return c.SendStatus(400)
	}
	secret := env.TokenSecret
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "Invalid request."})
	}
	//store req.body.password in a variable
	incoming := user.Password
	// Find user in database and return if hashed password matches, error if not matches or not exists
	query := bson.D{{Key: "email", Value: user.Email}}
	user = models.User{}
	if err := config.ConnDB("Users").FindOne(c.Context(), query).Decode(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "Invalid credentials."})
	}
	// Compare hashed password with incoming password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(incoming)); err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "Invalid credentials."})
	}
	// Create JWT
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
		"msg":   "User logged in successfully!",
		"token": jwt,
	})
}
