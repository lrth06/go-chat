package environment

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func HandleShutdown(app *fiber.App) {
	//TODO: Implement production shutdown
	fmt.Println("Running cleanup tasks...")
	// Cleanup Tasks
	fmt.Println("Removing tmp files...")
	//change ownership of tmp files
	os.Chmod("./tmp/", 0777)
	os.RemoveAll("./tmp/")
	os.RemoveAll("./logs/")
	fmt.Println("Cleanup tasks completed!")
	fmt.Println("Server Shutdown Successfully!")
	app.Shutdown()
	os.Exit(0)
}
