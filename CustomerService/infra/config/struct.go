package config

import "go.uber.org/zap"

type LoggerConfig struct {
	Config zap.Config
}

type DatabaseConfig struct {
	Driver   string
	Password string
	User     string
	Port     string
	Database string
}

type Config struct {
	Logger   LoggerConfig
	Database DatabaseConfig
}
