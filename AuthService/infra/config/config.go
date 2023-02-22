package config

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

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
