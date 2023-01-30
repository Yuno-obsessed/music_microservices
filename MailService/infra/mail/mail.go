package mail

import "gopkg.in/gomail.v2"

type Mail struct {
	From   string
	Dialer *gomail.Dialer
}

func (m *Mail) SendMail(to, subject, body string) error {
	message := gomail.NewMessage()

	message.SetHeader("From", m.From)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	return m.Dialer.DialAndSend(message)
}
