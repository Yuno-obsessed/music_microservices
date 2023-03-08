package config

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func NewSmtpServer() (*gomail.Dialer, error) {
	conf := SmtpConfigInit()
	dialer := gomail.NewDialer(conf.Host,
		conf.Port, conf.Username, conf.Password)
	conn, err := dialer.Dial()
	if err != nil {
		return nil, fmt.Errorf("error connecting to dialer, %v", err)
	}

	err = conn.Close()
	if err != nil {
		return nil, fmt.Errorf("error closing connection, %v", err)
	}
	return dialer, nil
}

type Smtp struct {
	Host     string
	Port     int
	Username string
	Password string
}

func SmtpConfigInit() Smtp {
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	return Smtp{
		Host:     os.Getenv("MAIL_HOST"),
		Port:     port,
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
	}
}
