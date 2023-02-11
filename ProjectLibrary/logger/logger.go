package logger

import (
	"go.uber.org/zap"
	"log"
	"project_library/config"
)

type Logger *zap.Logger

func NewLogger() Logger {
	conf := config.NewConfig()
	logger, err := conf.Logger.Config.Build()
	if err != nil {
		log.Fatalf("error setting up logger, %v", err)
	}
	return logger
}
