package logger

import (
	"database-service/config"
	"go.uber.org/zap"
	"log"
)

type Logger struct {
	Log *zap.Logger
}

func NewLogger(conf config.Config) Logger {
	logger, err := conf.Logger.Config.Build()
	if err != nil {
		log.Fatalf("error setting up logger, %v", err)
	}
	return Logger{
		Log: logger,
	}
}
