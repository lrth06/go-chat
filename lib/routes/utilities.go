package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/handlers"
	"github.com/lrth06/go-chat/lib/handlers/upload"
	"github.com/lrth06/go-chat/lib/middleware/token"
)

func RegisterUtilityRoutes(v fiber.Router) {
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

}
