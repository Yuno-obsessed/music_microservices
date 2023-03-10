package google

import (
	"os"

	"github.com/Yuno-obsessed/music_microservices/AuthService/infra/config"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func Google() {
	config.GothicConf()
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_APP_ID"),
			os.Getenv("GOOGLE_APP_SECRET"),
			"http://localhost:8081/api/v1/login/auth/google/callback"),
	)
}
