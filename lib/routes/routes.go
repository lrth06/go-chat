package routes

import (
	"github.com/gofiber/fiber/v2"
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

	RegisterUtilityRoutes(v)
	RegisterAuthRoutes(v)
	RegisterUserRoutes(v)
	RegisterRoomRoutes(v)

	//404 Wildcard (redirects to client which will route with react router)
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("client/build/index.html")
	})

}
