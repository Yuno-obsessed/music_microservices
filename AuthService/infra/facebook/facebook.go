package facebook

import (
	"os"

	"github.com/Yuno-obsessed/music_microservices/AuthService/infra/config"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/facebook"
)

func Facebook() {
	config.GothicConf()
	goth.UseProviders(
		facebook.New(os.Getenv("FACEBOOK_APP_ID"),
			os.Getenv("FACEBOOK_APP_SECRET"),
			"http://localhost:8081/api/v1/login/auth/facebook/callback"),
	)
}
