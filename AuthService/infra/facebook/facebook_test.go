package facebook_test

import (
	"auth-service/infra/facebook"
	"testing"

	"github.com/markbates/goth"
)

func TestRegisterApp(t *testing.T) {
	facebook.Facebook()
	providers := goth.GetProviders()
	if providers["facebook"] == nil {
		t.Errorf("Want providers, got %v", providers["facebook"].Name())
	}
}
