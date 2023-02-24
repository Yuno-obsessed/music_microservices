package config

import (
	"os"
	"strconv"
)

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
