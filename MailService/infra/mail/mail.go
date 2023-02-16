package mail

import (
	"github.com/Yuno-obsessed/music_microservices/MailService/domain/entity"
	"gopkg.in/gomail.v2"
)

type Mail struct {
	Dialer *gomail.Dialer
}

func NewMail() Mail {
	server, _ := NewSmtpServer()
	return server
}

func (m *Mail) SendMail(mail entity.Mail, msg MessageType) error {
	message := gomail.NewMessage()

	message.SetHeader("From", m.Dialer.Username)
	message.SetHeader("To", mail.Email)
	message.SetHeader("Subject", mail.Subject)
	message.SetBody("text/html", msg.Text())

	return m.Dialer.DialAndSend(message)
}
