package mail

import (
	"encoding/json"
	"fmt"
	"github.com/Yuno-obsessed/music_microservices/MailService/domain/entity"
	"github.com/Yuno-obsessed/music_microservices/MailService/infra/mail"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/gin-gonic/gin"
)

type Mailing struct {
	Mail   mail.Mail
	Logger logger.CustomLogger
}

func NewMailing() Mailing {
	return Mailing{mail.NewMail(),
		logger.NewLogger()}
}

// MailHandler is taking a payload of json "msg","recipient","subject"
func (m Mailing) MailSuccessfulRegistration(c *gin.Context) {
	mType := c.Param("message_type")
	var msg mail.MessageType
	switch mType {
	case mail.SuccessfulRegistration.Text():
		msg = mail.SuccessfulRegistration
		break
	case mail.SuccessfulLogin.Text():
		msg = mail.SuccessfulLogin
		break
	case mail.NewEventFromSubscriptions.Text():
		msg = mail.NewEventFromSubscriptions
		break
	default:
		m.Logger.Error("wrong endpoint calling mail service")
	}
	var newMail entity.Mail
	err := json.NewDecoder(c.Request.Body).Decode(&newMail)
	if err != nil {
		m.Logger.Error(fmt.Sprintf("Error sending mail, %v", err))
		c.AbortWithStatusJSON(400, fmt.Sprintf("error processing payload, %v", err))
	}
	err = m.Mail.SendMail(newMail, msg)
	if err != nil {
		m.Logger.Error(fmt.Sprintf("Error sending mail, %v", err))
	}
}
