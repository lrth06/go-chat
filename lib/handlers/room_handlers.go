package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/models"
	"github.com/lrth06/go-chat/lib/utils/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	return c.JSON(room)
}

func UpdateRoom(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.SendStatus(400)
	}
	room := models.Room{}
	if err := c.BodyParser(&room); err != nil {
		return c.SendStatus(400)
	}
	query := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: room}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	if err := config.ConnDB("Rooms").FindOneAndUpdate(c.Context(), query, update, opts).Decode(&room); err != nil {
		return c.SendStatus(400)
	}

	return c.JSON(fiber.Map{
		"msg":  "Room updated successfully!",
		"room": room,
	})
}

func DeleteRoom(c *fiber.Ctx) error {
	idParam := c.Params("id")
	roomId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "Invalid ID."})
	}
	query := bson.D{{Key: "_id", Value: roomId}}
	room := models.Room{}
	if err := config.ConnDB("Rooms").FindOne(c.Context(), query).Decode(&room); err != nil {
		return c.Status(404).JSON(fiber.Map{"msg": "Room not found."})
	}
	if _, err := config.ConnDB("Rooms").DeleteOne(c.Context(), query); err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "Room not deleted."})
	}
	// TODO: fix this (convert to production delete)
	//delete users images directory
	return c.JSON(fiber.Map{
		"msg": "Room deleted successfully!",
	})
}

func GetRooms(c *fiber.Ctx) error {
	query := bson.D{{}}
	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := config.ConnDB("Rooms").Find(c.Context(), query, opts)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var rooms []models.Room = make([]models.Room, 0)
	if err := cursor.All(c.Context(), &rooms); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(rooms)
}
