package structs

import (
	"github.com/gofiber/websocket/v2"
)

type Message struct {
	Data []byte `json:"data"`
	Room string `json:"room"`
	User string `json:"user"`
}

type Subscription struct {
	Conn *websocket.Conn `json:"conn"`
	Room string          `json:"room"`
	User string          `json:"user"`
}
