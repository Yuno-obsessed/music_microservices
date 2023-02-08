package config

import "go.uber.org/zap"

type Smtp struct {
	Host     string
	Port     int
	Username string
	Password string
}

type LoggerConfig struct {
	Config zap.Config
}

type Config struct {
	Smtp   Smtp
	Logger LoggerConfig
}
