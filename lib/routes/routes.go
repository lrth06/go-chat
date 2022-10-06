package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
	"github.com/lrth06/go-chat/lib/handlers"
	"github.com/lrth06/go-chat/lib/handlers/auth"
	"github.com/lrth06/go-chat/lib/handlers/rooms"
	"github.com/lrth06/go-chat/lib/handlers/upload"
	"github.com/lrth06/go-chat/lib/handlers/users"
	"github.com/lrth06/go-chat/lib/middleware/logging"
	"github.com/lrth06/go-chat/lib/middleware/token"
	"github.com/lrth06/go-chat/lib/utils/config"
)

func SetupRoutes(app *fiber.App) {

	env, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	appEnv := env.AppEnv

	app.Static("/", "./client/build")

	imagePath := "./client/build/images"

	if appEnv != "production" {
		imagePath = "./tmp/uploads/"
	}

	app.Static("/images", imagePath)

	app.Use(logging.Logger)

	app.Use(cors.New())

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(recover.New(
		recover.Config{
			EnableStackTrace: true,
		}))

	app.Use(etag.New())

	ws := app.Group("/ws", handlers.HandleUpgrade)
	//ws/room/b9fe28f7-9180-40e6-9488-36830507f7e1
	ws.Get("/room/:id", websocket.New(func(c *websocket.Conn) {
		handlers.HandleSocket(c)
	}))

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

	//user routes
	//api/v1/users
	v.Get("/users", users.GetUsers)

	//api/v1/user/
	user := v.Group("/user", func(c *fiber.Ctx) error {
		c.Set("API", "User")
		return c.Next()
	})
	UserRoutes(user)

	// //room routes
	// //api/v1/rooms
	v.Get("/rooms", rooms.GetRooms)

	// //api/v1/room/
	room := v.Group("/room", func(c *fiber.Ctx) error {
		c.Set("API", "Room")
		// RoomRoutes(v,c)

		return c.Next()
	})

		//use RoomRoutes subrouter
	RoomRoutes(room)


	//404 Wildcard (redirects to client which will route with react router)
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("client/build/index.html")
	})

}
