package mail

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"mail-service/infra/config"
)

func NewSmtpServer() (Mail, error) {
	conf := config.NewConfig()
	mail := Mail{
		Dialer: gomail.NewDialer(conf.Smtp.Host,
			conf.Smtp.Port, conf.Smtp.Username, conf.Smtp.Password),
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
