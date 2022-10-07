package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/handlers/rooms"
	"github.com/lrth06/go-chat/lib/middleware/permissions"
	"github.com/lrth06/go-chat/lib/middleware/token"
)

// take route group in and return a new route group
func RegisterRoomRoutes(v fiber.Router) {
	v.Get("/rooms", rooms.GetRooms)

	room := v.Group("/room", func(c *fiber.Ctx) error {
		c.Set("API", "Room")
		// RoomRoutes(v,c)

		return c.Next()
	})

	room.Post("/",
		token.ExtractToken,
			rooms.CreateRoom,
	)

	//api/v1/room/:id
	room.Patch("/:id",
		token.ExtractToken,
		permissions.AdminCheck,
		permissions.OwnerModerator,
			rooms.UpdateRoom,
	)
	room.Put("/:id",
		token.ExtractToken,
		permissions.AdminCheck,
		permissions.OwnerModerator,
			rooms.UpdateRoom,
	)

	//api/v1/room/:id
	room.Get("/:id", rooms.GetRoom)

	//api/v1/room/:id
	room.Delete("/:id",
		token.ExtractToken,
		permissions.AdminCheck,
		permissions.OwnerModerator,
			rooms.DeleteRoom,
	)

}
