package mail

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"mail-service/domain/entity"
	"mail-service/infra/logger"
	"mail-service/infra/mail"
)

type Mailing struct {
	Mail   mail.Mail
	Logger logger.Logger
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
		m.Logger.Log.Error("wrong endpoint calling mail service")
	}
	var newMail entity.Mail
	err := json.NewDecoder(c.Request.Body).Decode(&newMail)
	if err != nil {
		m.Logger.Log.Error(fmt.Sprintf("Error sending mail, %v", err))
		c.AbortWithStatusJSON(400, fmt.Sprintf("error processing payload, %v", err))
	}
	err = m.Mail.SendMail(newMail, msg)
	if err != nil {
		m.Logger.Log.Error(fmt.Sprintf("Error sending mail, %v", err))
	}
}

// TODO: mailhandler, migrations for authService, think of roles implementing
// TODO: CustomerService structure
