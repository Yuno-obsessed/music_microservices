package mail_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Yuno-obsessed/music_microservices/MailService/domain/entity"
	"github.com/Yuno-obsessed/music_microservices/MailService/infra/server/handlers/mail"
	"github.com/gin-gonic/gin"
)

func TestMailing_MailSuccessfulRegistration(t *testing.T) {
	mock := entity.Mail{
		"Successful registration",
		"d.2510086@gmail.com",
		"s.danilo2406@gmail.com",
		"registration was completed with success",
	}
	payload, err := json.Marshal(mock)
	if err != nil {
		t.Errorf("error marshalling mocks, %v", err)
	}
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request, err = http.NewRequest("POST", "http://localhost:8082/api/v1/mail-service/send", bytes.NewBuffer(payload))
	if err != nil {
		t.Errorf("error in request, %v", err)
	}

	mailing := mail.NewMailing()
	mailing.MailSuccessfulRegistration(c)
}
