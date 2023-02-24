package mail

import (
	"encoding/json"
	"fmt"

	"github.com/Yuno-obsessed/music_microservices/MailService/domain/entity"
	"github.com/Yuno-obsessed/music_microservices/MailService/infra/mail"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	mtypes "github.com/Yuno-obsessed/music_microservices/ProjectLibrary/mail"
	"github.com/gin-gonic/gin"
)

// Pass here a service
type Mailing struct {
	Mail   mail.Mail
	Logger logger.CustomLogger
}

func NewMailing() Mailing {
	return Mailing{
		mail.NewMail(),
		logger.NewLogger(),
	}
}

// MailHandler is taking a payload of json "msg","recipient","subject"
func (m Mailing) MailSuccessfulRegistration(c *gin.Context) {
	mType := c.Param("message_type")
	// how do I call it the best way as an api? take mtype from headers or json payload?
	var msg mtypes.MessageType
	switch mType {
	case mtypes.SuccessfulRegistration.Text():
		msg = mtypes.SuccessfulRegistration
		break
	case mtypes.SuccessfulLogin.Text():
		msg = mtypes.SuccessfulLogin
		break
	case mtypes.NewEventFromSubscriptions.Text():
		msg = mtypes.NewEventFromSubscriptions
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
	err = m.Mail.SendMail(newMail)
	if err != nil {
		m.Logger.Error(fmt.Sprintf("Error sending mail, %v", err))
	}
}
