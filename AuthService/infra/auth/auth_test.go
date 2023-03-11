package auth_test

import (
	"testing"

	"github.com/Yuno-obsessed/music_microservices/AuthService/infra/auth"

	"github.com/kataras/jwt"
)

func TestJWT_GenerateToken(t *testing.T) {
	jwtStruct := auth.NewJWT()
	username := "yuno"
	token, err := jwtStruct.GenerateToken(username)
	if err != nil {
		t.Errorf("Error generating token, got %v", token)
	}
	verified, err := jwt.Verify(jwt.HS256, jwtStruct.Secret, []byte(token))
	if err != nil {
		t.Errorf("Error validating token, expects: %v, got: %v", token, verified.Token)
	}
}
