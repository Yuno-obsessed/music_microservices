package config

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func NewConfig() Config {
	return Config{
		Logger:   LoggerConfigInit(),
		Database: DatabaseConfigInit(),
	}
}

func DatabaseConfigInit() DatabaseConfig {
	godotenv.Load("../../.env")
	return DatabaseConfig{
		Driver:   os.Getenv("POSTGRES_DRIVER"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DB"),
	}
}

func LoggerConfigInit() LoggerConfig {
	godotenv.Load("../../.env")
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
