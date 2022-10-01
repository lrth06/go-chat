// import server from main and run unit tests
package main

import (
	"net/http"
	"testing"

	"github.com/lrth06/go-chat/lib/structs"
)

func TestTruthiness(t *testing.T) {
	config := structs.Config{
		Port:   "3000",
		AppEnv: "test",
	}
	app := Server(config)
	if app == nil {
		t.Error("app is nil")
	}

	//test public routes

	routes := []string{
		"/api",
	}

	for _, route := range routes {
		req, err := http.NewRequest("GET", route, nil)
		if err != nil {
			t.Error(err)
		}
		resp, err := app.Test(req)
		if err != nil {
			t.Error(err)
		}
		if resp.StatusCode != 200 {
			t.Errorf("Expected 200, got %d", resp.StatusCode)
		}

	}

}
