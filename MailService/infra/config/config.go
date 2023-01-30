package config

import (
	"os"
	"strconv"
)

func NewConfig() Config {
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	return Config{
		Smtp: Smtp{
			Host:     os.Getenv("MAIL_HOST"),
			Port:     port,
			Username: os.Getenv("MAIL_USERNAME"),
			Password: os.Getenv("MAIL_PASSWORD"),
		},
	}
}
