package config

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"strconv"
)

func NewConfig() Config {
	godotenv.Load("../../.env")
	return Config{
		Smtp:   SmtpConfigInit(),
		Logger: LoggerConfigInit(),
	}
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
func LoggerConfigInit() LoggerConfig {
	return LoggerConfig{
		Config: zap.Config{
			Level:            zap.NewAtomicLevel(),
			Development:      true,
			Encoding:         "json",
			EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
			OutputPaths:      []string{"../../logs/info.log"},
			ErrorOutputPaths: []string{"../../logs/error.log"},
		},
	}
}
