package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/lrth06/go-chat/lib/handlers"
)

func RegisterWebsocketRoutes(app *fiber.App) {
	ws := app.Group("/ws", handlers.HandleUpgrade)
	//ws/room/b9fe28f7-9180-40e6-9488-36830507f7e1
	ws.Get("/room/:id", websocket.New(func(c *websocket.Conn) {
		handlers.HandleSocket(c)
	}))

}
