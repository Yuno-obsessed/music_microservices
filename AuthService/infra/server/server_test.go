package server_test

import (
	"auth-service/infra/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Searching for ideas about useful tests for router

func TestInit(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal()
	}

	rr := httptest.NewRecorder()
	router := server.NewRouter()
	router.InitRoutes()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != 200 {
		t.Errorf("Init() returned wrong status code: got %v want %v", status, 200)
	}

	routes := router.Routes()
	i := 0
	r := routes[i]
	for r.HandlerFunc != nil {
		if r.Method != "POST" && r.Path != "/api/v1/login" {
			t.Errorf("Init() did not initialize the /api/v1/login POST route")
		}
		if r.Method != "GET" && r.Path != "/api/v1/login/auth/facebook" {
			t.Errorf("Init() did not initialize the /api/v1/login/auth/facebook GET route")
		}
		if r.Method != "GET" && r.Path != "/api/v1/login/auth/facebook/callback" {
			t.Errorf("Init() did not initialize the /api/v1/login/auth/facebook/callback GET route")
		}
		i++
	}
}
