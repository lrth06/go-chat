package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/lrth06/go-chat/lib/structs"
)

func HandleStartup(config structs.Config) {
	if config.AppEnv == "development" {
		os.MkdirAll("./tmp/uploads/users", os.ModePerm)
		os.MkdirAll("./tmp/server", os.ModePerm)
		os.MkdirAll("./tmp/logs", os.ModePerm)
		//create file named access-log.json in tmp/logs folder
		logs, err := os.OpenFile("./tmp/logs/access-log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		// msg := "Running in " + env + " Environment at " + time.Now().Format("2006-01-02 15:04:05")
		// LogItem("INFO", msg)
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
