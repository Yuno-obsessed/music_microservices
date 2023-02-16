package mail

import (
	"fmt"
	"github.com/Yuno-obsessed/music_microservices/MailService/infra/config"
	"gopkg.in/gomail.v2"
)

func NewSmtpServer() (Mail, error) {
	conf := config.SmtpConfigInit()
	mail := Mail{
		Dialer: gomail.NewDialer(conf.Host,
			conf.Port, conf.Username, conf.Password),
	}
	conn, err := mail.Dialer.Dial()
	if err != nil {
		return Mail{}, fmt.Errorf("error connecting to dialer, %v", err)
	}

	err = conn.Close()
	if err != nil {
		return Mail{}, fmt.Errorf("error closing connection, %v", err)
	}
	return mail, nil
}
