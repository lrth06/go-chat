package logging

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lrth06/go-chat/lib/structs"
	"github.com/lrth06/go-chat/lib/utils/config"
)

func Logger(c *fiber.Ctx) error {
	// beautify the request body
	var body interface{}
	err := json.Unmarshal(c.Body(), &body)
	if err != nil {
		body = string(c.Body())
	}

	payload := structs.Payload{
		Status:  c.Response().StatusCode(),
		Method:  c.Method(),
		Path:    c.Path(),
		IP:      c.IP(),
		Headers: c.GetReqHeaders(),
		Body:    body,
	}

	LogItem("INFO", payload)
	return c.Next()
}

//HACK This needs to be refactored ASAP

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
