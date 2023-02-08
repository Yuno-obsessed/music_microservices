package logger

import (
	"go.uber.org/zap"
	"log"
	"mail-service/infra/config"
)

type Logger struct {
	Log *zap.Logger
}

func NewLogger() Logger {
	conf := config.NewConfig()
	logger, err := conf.Logger.Config.Build()
	if err != nil {
		log.Fatalf("error setting up logger, %v", err)
	}
	return Logger{
		Log: logger,
	}
}
