package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/lrth06/go-chat/lib/structs"
)

func HandleStartup(config structs.Config) {
	// TODO: Implement production startup
	if config.AppEnv == "development" {
		os.MkdirAll("./tmp/uploads/users", os.ModePerm)
		os.MkdirAll("./tmp/server", os.ModePerm)
		os.MkdirAll("./tmp/logs", os.ModePerm)
		logs, err := os.OpenFile("./tmp/logs/access-log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer logs.Close()
		f, err := os.Create("./tmp/server/pid")
		if err != nil {
			panic(err)
		}
		f.WriteString(fmt.Sprintf("%d", os.Getpid()))
		defer f.Close()
		return
	}
}
