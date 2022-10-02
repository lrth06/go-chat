// import server from main and run unit tests
package main

import (
	"net/http"
	"testing"

	"github.com/lrth06/go-chat/lib/utils/config"
)

func TestTruthiness(t *testing.T) {
	env, err := config.GetConfig()
	if err != nil {
		t.Error(err)
	}
	app := Server(env)
	if app == nil {
		t.Error("app is nil")
	}

	routes := []string{
		"/",
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

	//log success and exit with code 0
	t.Log("Test passed")

}

func TestMaths(t *testing.T) {
	if 1+1 != 2 {
		t.Error("1 + 1 should equal 2")
	}
}
