package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/handlers/users"
	"github.com/lrth06/go-chat/lib/middleware/permissions"
	"github.com/lrth06/go-chat/lib/middleware/token"
	"github.com/lrth06/go-chat/lib/middleware/validation"
)

// take route group in and return a new route group
func RegisterUserRoutes(v fiber.Router) {
	v.Get("/users", users.GetUsers)

	//api/v1/user/
	user := v.Group("/user", func(c *fiber.Ctx) error {
		c.Set("API", "User")
		return c.Next()
	})

	user.Post("/",
	validation.ValidateUser,
		users.CreateUser,
	)
	//api/v1/user/:id
	user.Get("/:id",
		token.ExtractToken,
		permissions.AdminCheck,
		permissions.SelfCheck,
			users.GetUser,
	)

	//api/v1/user/:id
	user.Patch("/:id",
		token.ExtractToken,
		permissions.AdminCheck,
		permissions.SelfCheck,
			users.UpdateUser,
	)
	user.Put("/:id",
		token.ExtractToken,
		permissions.AdminCheck,
		permissions.SelfCheck,
			users.UpdateUser,
	)

	//api/v1/user/:id
	user.Delete("/:id",
		token.ExtractToken,
		permissions.SelfCheck,
		permissions.AdminCheck,
			users.DeleteUser,
)

}
