package structs

import (
	"fmt"

	"github.com/segmentio/encoding/json"
)


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
	Body    string            `json:"body"`
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
