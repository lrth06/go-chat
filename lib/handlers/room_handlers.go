package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/models"
	"github.com/lrth06/go-chat/lib/utils/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: Implement this entire file

func CreateRoom(c *fiber.Ctx) error {
	//assign to room model
	room := models.Room{}
	room.ID = primitive.NewObjectID()
	if err := c.BodyParser(&room); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": "Invalid request.",
		})
	}
	fmt.Println(room)
	//set Owner to the user who created the room by c.Get("id")
	//get primitive.ObjectID from string in c.Locals("id")
	room.Owner, _ = primitive.ObjectIDFromHex(c.Locals("id").(string))

	//configure room.Settings
	//reject if required fields are empty
	if room.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"msg": "Room name is required!",
		})
	}
	//check for duplicate room name
	filter := bson.M{
		"name": room.Name,
	}
	if err := config.ConnDB("Rooms").FindOne(c.Context(), filter).Decode(&room); err == nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": "Room already exists!",
		})
	}
	//insert room into database
	if _, err := config.ConnDB("Rooms").InsertOne(c.Context(), room); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(room)

}

func GetRoom(c *fiber.Ctx) error {
	return c.Status(500).JSON(map[string]string{
		"msg": "Not implemented",
	})
}

func UpdateRoom(c *fiber.Ctx) error {
	return c.Status(500).JSON(map[string]string{
		"msg": "Not implemented",
	})
}

func DeleteRoom(c *fiber.Ctx) error {
	return c.Status(500).JSON(map[string]string{
		"msg": "Not implemented",
	})
}

func GetRooms(c *fiber.Ctx) error {
	query := bson.D{{}}
	cursor, err := config.ConnDB("Rooms").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var rooms []models.Room = make([]models.Room, 0)
	if err := cursor.All(c.Context(), &rooms); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(rooms)
}
