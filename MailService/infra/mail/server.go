package mail

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"mail-service/infra/config"
)

func NewSmtpServer(conf config.Config) (Mail, error) {
	mail := Mail{
		Dialer: gomail.NewDialer(conf.Smtp.Host,
			conf.Smtp.Port, conf.Smtp.Username, conf.Smtp.Password),
		From: conf.Smtp.Username,
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
