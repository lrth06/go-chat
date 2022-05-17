package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func HandleShutdown(app *fiber.App, c chan os.Signal) {
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	//lint:ignore S1005 This is a simple way to wait for a signal
	_ = <-c
	fmt.Println("\nGracefully shutting down...")
	_ = app.Shutdown()
	fmt.Println("Running cleanup tasks...")
	// Cleanup Tasks
	fmt.Println("Removing tmp files...")
	//change ownership of tmp files
	os.Chmod("./tmp/", 0777)
	os.RemoveAll("./tmp/")
	os.RemoveAll("./logs/")
	close(c)
	fmt.Println("Cleanup tasks completed!")
	fmt.Println("Server Shutdown Successfully!")
	os.Exit(0)
}
