package config

import (
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"go.uber.org/zap"
)

type Logger *zap.Logger

func NewLogger() Logger {
	return logger.NewLogger()
}
