package logger

import (
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/config"
	"go.uber.org/zap"
	"log"
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
