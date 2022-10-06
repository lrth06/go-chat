package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/handlers/users"
	"github.com/lrth06/go-chat/lib/middleware/permissions"
	"github.com/lrth06/go-chat/lib/middleware/token"
	"github.com/lrth06/go-chat/lib/middleware/validation"
)

// take route group in and return a new route group
func UserRoutes(user fiber.Router) {
	user.Post("/",
	validation.ValidateUser,
	users.CreateUser,
)
//api/v1/user/:id
user.Get("/:id", users.GetUser)

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
	permissions.AdminCheck,
	users.DeleteUser,
)

}
