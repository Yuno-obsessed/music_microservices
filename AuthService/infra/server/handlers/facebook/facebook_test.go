package facebook_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Yuno-obsessed/music_microservices/AuthService/infra/server"

	"github.com/joho/godotenv"
	"github.com/markbates/goth"
)

var (
	rr     *httptest.ResponseRecorder
	router server.Router
)

func setup() {
	godotenv.Load("../../../.env")
	router = server.NewRouter()
	rr = httptest.NewRecorder()
}

func TestFacebookLoginHandler(t *testing.T) {
	setup()
	// godotenv.Load("../../../.env")
	req, err := http.NewRequest("GET", "http://localhost:8081/api/v1/login/auth/facebook", nil)
	if err != nil {
		t.Errorf("Error in request, %v", err)
	}

	router.ServeHTTP(rr, req)

	if rr.Code == 200 {
		t.Fatalf("Handler return the wrong status: got %v, want %v", rr.Code, 200)
	}
}

func TestFacebookLoginCallback(t *testing.T) {
	setup()
	// godotenv.Load("../../../.env")
	req, err := http.NewRequest("GET", "http://localhost:8081/api/v1/login/auth/facebook/callback", nil)
	if err != nil {
		t.Errorf("Error in request, %v", err)
	}

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != 200 {
		t.Errorf("Handler returned the wrong status: got %v, want %v", rr.Body.String(), 200)
	}

	var gothUser goth.User
	got := json.NewDecoder(rr.Body).Decode(&gothUser)
	if got == nil {
		t.Errorf("Expected to return goth.User struct, got %v", gothUser)
	}
}
