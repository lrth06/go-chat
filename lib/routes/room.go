package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/handlers/rooms"
	"github.com/lrth06/go-chat/lib/middleware/permissions"
	"github.com/lrth06/go-chat/lib/middleware/token"
)

// take route group in and return a new route group
func RoomRoutes(room fiber.Router) {

	room.Post("/",
		token.ExtractToken,
		rooms.CreateRoom,
	)

	//api/v1/room/:id
	room.Patch("/:id",
		permissions.AdminCheck,
		rooms.UpdateRoom,
	)
	room.Put("/:id",
		permissions.AdminCheck,
		rooms.UpdateRoom,
	)

	//api/v1/room/:id
	room.Get("/:id", rooms.GetRoom)

	//api/v1/room/:id
	room.Delete("/:id",
		permissions.AdminCheck,
		rooms.DeleteRoom,
	)

}
