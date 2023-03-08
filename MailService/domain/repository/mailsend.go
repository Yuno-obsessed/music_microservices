package repository

import (
	"github.com/Yuno-obsessed/music_microservices/MailService/domain/entity"
	"github.com/Yuno-obsessed/music_microservices/MailService/infra/config"
	"gopkg.in/gomail.v2"
)

type MailSend struct {
	Dialer *gomail.Dialer
}

func NewMail() *MailSend {
	server, _ := config.NewSmtpServer()
	return &MailSend{
		server,
	}
}

func (m *MailSend) SendMail(mail entity.Mail) error {
	message := gomail.NewMessage()

	message.SetHeader("From", m.Dialer.Username)
	message.SetHeader("To", mail.Email)
	message.SetHeader("Subject", mail.Subject)
	message.SetBody("text/html", mail.Type.Text())

	return m.Dialer.DialAndSend(message)
}

// In the microservice where getting qr code for ticket, to send a request to upload service
// first to get a generated qr code there and then to call this method with mail?
func (m *MailSend) SendMailWithAttachment(mail entity.Mail) error {
	message := gomail.NewMessage()

	return m.Dialer.DialAndSend(message)
}
