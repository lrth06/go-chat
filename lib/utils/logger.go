package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lrth06/go-chat/lib/structs"
	"github.com/lrth06/go-chat/lib/utils/config"
)

func LogItem(logLevel string, payload structs.Payload) error {
	env, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	appEnv := env.AppEnv

	var logItem structs.LogItem
	logItem.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	logItem.Level = logLevel
	logItem.ID = uuid.New().String()
	logItem.Payload = payload

	//unmarshal the log item into a json string
	jsonString, err := json.Marshal(logItem)
	if err != nil {
		fmt.Println(err)
	}

	if appEnv == "development" {
		logs, err := os.OpenFile("./tmp/logs/access-log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer logs.Close()
		//indent the json string for readability
		jsonString, err = json.MarshalIndent(logItem, "", "    ")
		if err != nil {
			fmt.Println(err)
		}
		//indent jsonString.payload for readability
		//parse json string and indent the payload for readability
		var logItem structs.LogItem
		err = json.Unmarshal([]byte(jsonString), &logItem)
		if err != nil {
			fmt.Println(err)
		}
		jsonString, err = json.MarshalIndent(logItem, "", "    ")
		if err != nil {
			fmt.Println(err)
		}
		// remove all \ from the json string
		jsonString = []byte(strings.ReplaceAll(string(jsonString), "\\", ""))
		_, err = logs.WriteString(string(jsonString) + ",\n")
		if err != nil {
			fmt.Println(err)
		}
	}
	println(string(jsonString))

	return err
}
