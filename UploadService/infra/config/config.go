package config

import (
	"go.uber.org/zap"
	"project_library/logger"
)

type Logger *zap.Logger

func NewLogger() Logger {
	return logger.NewLogger()
}
