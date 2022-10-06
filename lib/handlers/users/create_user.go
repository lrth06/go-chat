package users

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/lrth06/go-chat/lib/models"
	"github.com/lrth06/go-chat/lib/utils/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	env, err := config.GetConfig()
	if err != nil {
		return c.SendStatus(400)
	}
	secret := env.TokenSecret
	//temporary struct to store incoming confirmation password
	type passVerification struct {
		Password2 string `json:"password2"`
	}

	user := models.User{}

	pass2 := passVerification{}
	if err := c.BodyParser(&user); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{"msg": "Invalid request."})
	}
	if err := c.BodyParser(&pass2); err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "Password confirmation is required!"})
	}
	// reject if required fields are empty
	if user.Name == "" || user.Password == "" || user.Email == "" {
		return c.Status(400).JSON(fiber.Map{"msg": "Username, password and email are required!"})
	}
	// [x] Validate HERE

	if user.Password != pass2.Password2 {
		return c.Status(400).JSON(fiber.Map{
			"msg": "Passwords do not match!",
		})
	}
	//set filter to be mongo OR query for email OR name
	filter := bson.M{
		"$or": bson.A{
			bson.M{"email": user.Email},
			bson.M{"name": user.Name},
		},
	}
	if err := config.ConnDB("Users").FindOne(c.Context(), filter).Decode(&user); err == nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": "Email/Username already exists!",
		})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.SendStatus(400)
	}
	user.Password = string(hashedPassword)
	user.ID = primitive.NewObjectID()
	user.Avatar = "https://avatars.dicebear.com/api/bottts/" + user.Name + ".svg"
	user.Token = base64.StdEncoding.EncodeToString([]byte(user.ID.Hex() + ":" + user.Email))
	encodedRole := base64.StdEncoding.EncodeToString([]byte("user:basic"))
	user.Roles = []string{encodedRole}
	if _, err := config.ConnDB("Users").InsertOne(c.Context(), user); err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "User not created."})
	}
	jwtClaims := jwt.MapClaims{
		"id":     user.ID,
		"email":  user.Email,
		"name":   user.Name,
		"avatar": user.Avatar,
		"roles":  user.Roles,
		"exp":    jwt.TimeFunc().Add(time.Hour * 24).Unix(),
	}

	jwt, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims).SignedString([]byte(secret))
	if err != nil {
		return c.SendStatus(400)
	}
	return c.JSON(fiber.Map{
		"msg":   "User created successfully!",
		"token": jwt,
	})
}
