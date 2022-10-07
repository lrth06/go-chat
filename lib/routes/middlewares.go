package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/lrth06/go-chat/lib/middleware/logging"
	"github.com/lrth06/go-chat/lib/utils/config"
)

func RegisterMiddlewares(app *fiber.App) {
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
			EnableStackTrace: env.AppEnv != "production",
		}))

	app.Use(etag.New())
}
