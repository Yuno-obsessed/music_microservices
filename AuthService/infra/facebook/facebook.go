package facebook

import (
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/facebook"
	"os"
)

func Facebook() {
	goth.UseProviders(
		facebook.New(os.Getenv("FACEBOOK_APP_ID"),
			os.Getenv("FACEBOOK_APP_SECRET"),
			"http://localhost:8081/auth/facebook/callback"),
	)
}
