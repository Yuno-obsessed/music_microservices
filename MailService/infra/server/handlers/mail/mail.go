package mail

import (
	"encoding/json"
	"strings"

	"go.uber.org/zap"

	"github.com/Yuno-obsessed/music_microservices/MailService/domain/entity"
	"github.com/Yuno-obsessed/music_microservices/MailService/infra/mail"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/gin-gonic/gin"
)

// Pass here a service
type Mailing struct {
	Mail   mail.Mail
	Logger logger.CustomLogger
}

func NewMailing() *Mailing {
	return &Mailing{
		mail.NewMail(),
		logger.NewLogger(),
	}
}

// MailHandler is taking a payload of json "msg","recipient","subject"
func (m *Mailing) SendMail(c *gin.Context) {
	var newMail entity.Mail

	err := json.NewDecoder(c.Request.Body).Decode(&newMail)
	if err != nil {
		m.Logger.Error("error sending mail", zap.Error(err))
		c.AbortWithStatusJSON(400, gin.H{"error processing payload, %v": err})
	}
	err = m.Mail.SendMail(newMail)
	if err != nil {
		m.Logger.Error("error sending mail", zap.Error(err))
		c.AbortWithStatusJSON(500, gin.H{"error sending mail": err})
	}
}

func (m *Mailing) SendMailWithUpload(c *gin.Context) {
	cookie, err := c.Cookie("mail")
	if err != nil {
		m.Logger.Error("error processing cookie", zap.Error(err))
		c.AbortWithStatusJSON(500, gin.H{"error processing cookie": err})
	}
	cookieValues := strings.Split(cookie, ":")
}
