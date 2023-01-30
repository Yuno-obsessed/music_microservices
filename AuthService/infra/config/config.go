package config

import (
	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
	"os"
)

type DatabaseConfig struct {
	Driver   string
	Password string
	User     string
	Port     string
	Database string
}

func GothicConf() {
	key := os.Getenv("SESSION_SECRET")
	maxAge := 86400 * 30
	isProd := false // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store
}

func DatabaseConfigInit() *DatabaseConfig {
	return &DatabaseConfig{
		Driver:   os.Getenv("POSTGRES_DRIVER"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DB"),
	}
}
