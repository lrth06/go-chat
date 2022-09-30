// import server from main and run unit tests
package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/lrth06/go-chat/lib/utils/config"
)

func TestTruthiness(t *testing.T) {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := Server(config)
	if app == nil {
		t.Error("app is nil")
	}

	//test public routes

	routes := []string{
		"/",
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
