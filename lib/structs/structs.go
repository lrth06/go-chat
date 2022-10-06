package structs

import (
	"fmt"

	"github.com/segmentio/encoding/json"
)

type Config struct {
	Port        string `json:"port"`
	AppEnv      string `json:"app_env"`
	MongoURI    string `json:"mongo_uri"`
	DBName      string `json:"db_name"`
	TokenSecret string `json:"token_secret"`
}
type LogItem struct {
	Timestamp string  `json:"timestamp"`
	ID        string  `json:"id"`
	Level     string  `json:"level"`
	Payload   Payload `json:"payload"`
}

type Payload struct {
	Method  string            `json:"method"`
	Status  int               `json:"status"`
	Path    string            `json:"path"`
	IP      string            `json:"ip"`
	Headers map[string]string `json:"headers"`
}

func (p Payload) String() string {
	jsonString, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	//return unescaped string
	return string(jsonString)
}
func (p Payload) Raw() string {
	//convert to raw json string
	jsonString, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	raw := json.RawMessage(jsonString)
	return string(raw)
}

func (l LogItem) String() string {
	jsonString, err := json.Marshal(l)
	if err != nil {
		fmt.Println(err)
	}
	return string(jsonString)
}

type User struct {
	ID     any           `json:"_id"`
	Name   string        `json:"username"`
	Email  string        `json:"email" validate:"required,email"`
	Avatar string        `json:"avatar" validate:"omitempty,url"`
	Token  string        `json:"token"`
	Roles  []interface{} `json:"roles"`
}

type Room struct {
	Id    string `json:"id"`
	Users []User `json:"users"`
}
