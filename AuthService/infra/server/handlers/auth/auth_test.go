package auth_test

import (
	"auth-service/domain/entity"
	jwt "auth-service/infra/auth"
	"auth-service/infra/server/handlers/auth"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
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
