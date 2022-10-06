package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/handlers/auth"
)

func RegisterAuthRoutes(v fiber.Router) {

	//auth routes
	//api/v1/auth
	authRoutes := v.Group("/auth", func(c *fiber.Ctx) error {
		c.Set("API", "Auth")
		return c.Next()
	})
	//api/v1/auth/login
	authRoutes.Post("/login", auth.LoginUser)
	//api/v1/auth/logout
	authRoutes.Post("/logout", auth.LogoutUser)
}
