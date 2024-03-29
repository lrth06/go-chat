package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/routes"
	"github.com/lrth06/go-chat/lib/structs"
	"github.com/lrth06/go-chat/lib/utils/config"
	"github.com/lrth06/go-chat/lib/utils/environment"
)

const idleTimeout = 5 * time.Second

func main() {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := Server(config)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		//lint:ignore S1005 This is a simple way to wait for a signal
		_ = <-c
		//cleanup tasks here
		fmt.Println("\nShutting down server...")
		environment.HandleShutdown(app)
		_ = app.Shutdown()
	}()

	if err := app.Listen(":" + config.Port); err != nil {
		log.Panic(err)
	}
}

func Server(config structs.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  idleTimeout,
		// Prefork:               config.AppEnv == "production",
		AppName:               "Go-Chat",
		DisableStartupMessage: config.AppEnv == "production",
	})
	// app.Server().MaxConnsPerIP = 1

	environment.HandleStartup(config)
	routes.SetupRoutes(app)
	return app
}
