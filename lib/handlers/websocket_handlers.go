package handlers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/lrth06/go-chat/lib/structs"
)

type hub struct {
	rooms      map[string]map[*websocket.Conn]bool
	broadcast  chan structs.Message
	register   chan structs.Subscription
	unregister chan structs.Subscription
}

var h = hub{
	broadcast:  make(chan structs.Message),
	register:   make(chan structs.Subscription),
	unregister: make(chan structs.Subscription),
	rooms:      make(map[string]map[*websocket.Conn]bool),
}

func HandleUpgrade(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}



func (h *hub) Run() {
	for {
		select {
		case s := <-h.register:
			connections := h.rooms[s.Room]
			if connections == nil {
				connections = make(map[*websocket.Conn]bool)
				h.rooms[s.Room] = connections
			}
			h.rooms[s.Room][s.Conn] = true
		case s := <-h.unregister:
			connections := h.rooms[s.Room]
			if connections != nil {
				//lint:ignore S1033 This is a needed safety check to prevent a panic
				if _, ok := connections[s.Conn]; ok {
					delete(connections, s.Conn)
				}
			}
		case m := <-h.broadcast:
			connections := h.rooms[m.Room]
			for connection := range connections {
				if err := connection.WriteMessage(websocket.TextMessage, []byte(m.Data)); err != nil {
					delete(connections, connection)
					if len(connections) == 0 {
						delete(h.rooms, m.Room)
					}
				}
			}
		}
	}
}

func HandleSocket(c *websocket.Conn) {
	go h.Run()
	roomId := c.Params("id")
	user := c.Query("token")
	s := structs.Subscription{Conn: c, Room: roomId, User: user}
	defer func() {
		h.unregister <- s
		s.Conn.Close()
	}()
	h.register <- s
	for {
		messageType, msg, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("user disconnected", s.User)
				log.Println("read error: ", err)
			}
			return
		}
		if messageType == websocket.TextMessage {
			fmt.Println("message: ", string(msg))
			m := structs.Message{Data: msg, Room: s.Room}
			h.broadcast <- m
		} else {
			log.Println("websocket message received of type", messageType)
		}
	}
}
