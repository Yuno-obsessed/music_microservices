package config

import "go.uber.org/zap"

type DatabaseConfig struct {
	Driver   string
	Password string
	User     string
	Port     string
	Database string
}

type LoggerConfig struct {
	Config zap.Config
}

type Config struct {
	Database DatabaseConfig
	Logger   LoggerConfig
}
