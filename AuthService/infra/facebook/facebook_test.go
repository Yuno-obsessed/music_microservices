package facebook_test

import (
	"testing"

	"github.com/Yuno-obsessed/music_microservices/AuthService/infra/facebook"

	"github.com/markbates/goth"
)

func TestRegisterApp(t *testing.T) {
	facebook.Facebook()
	providers := goth.GetProviders()
	if providers["facebook"] == nil {
		t.Errorf("Want providers, got %v", providers["facebook"].Name())
	}
}
