package logger

import (
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/config"
	"go.uber.org/zap"
	"log"
)

type CustomLogger struct {
	*zap.Logger
}

func NewLogger() CustomLogger {
	conf := config.LoggerConfigInit()
	logger, err := conf.Config.Build()
	if err != nil {
		log.Fatalf("error setting up logger, %v", err)
	}
	return CustomLogger{logger}
}
