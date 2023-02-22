package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Yuno-obsessed/music_microservices/AuthService/domain/entity"
	jwt "github.com/Yuno-obsessed/music_microservices/AuthService/infra/auth"
	"github.com/Yuno-obsessed/music_microservices/AuthService/infra/server/handlers/auth"
	"github.com/gin-gonic/gin"
)

func TestLoginHandler(t *testing.T) {
	mock := entity.Login{
		"example@gmail.com",
		"wordpass",
	}
	payload, err := json.Marshal(mock)
	if err != nil {
		t.Errorf("error marshalling mocks, %v", err)
	}
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request, err = http.NewRequest("POST", "http://localhost:8081/login", bytes.NewBuffer(payload))
	if err != nil {
		t.Errorf("error in request, %v", err)
	}

	auth.LoginHandler(c)

	cookie := rr.Result().Cookies()[0]
	if cookie.Name != "Authorization" {
		t.Errorf("Authorization cookie wasn't found, got %v", cookie.Name)
	}
	token, _ := jwt.NewJWT().GenerateToken(mock.Email)
	if cookie.Value != token {
		t.Errorf("Expected token: %v, \n\t\tgot: %v", token, cookie.Value)
	}
}
