package auth_test

import (
	"auth-service/infra/auth"
	"github.com/kataras/jwt"
	"testing"
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
