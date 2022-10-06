package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/handlers"
	"github.com/lrth06/go-chat/lib/handlers/upload"
	"github.com/lrth06/go-chat/lib/middleware/token"
)

func SetupRoutes(app *fiber.App) {

	RegisterMiddlewares(app)
	RegisterWebsocketRoutes(app)

	api := app.Group("/api")

	//api/v1
	v := api.Group("/:version", func(c *fiber.Ctx) error {
		c.Set("Version", c.Params("version"))
		return c.Next()
	})

	v.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg": "Thank you for using the go-chat api, please refer to the documentation for more information.",
		})
	})

	//api/v1/random
	v.Get("/random", handlers.GetRandomID)
	//api/v1/random
	v.Post("/upload",
		token.ExtractToken,
		upload.HandleUpload,
	)

	RegisterAuthRoutes(v)
	RegisterUserRoutes(v)
	RegisterRoomRoutes(v)


	//404 Wildcard (redirects to client which will route with react router)
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("client/build/index.html")
	})

}
