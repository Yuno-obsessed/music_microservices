package security_test

import (
	"auth-service/infra/security"
	"testing"
)

func TestHash(t *testing.T) {
	password := "secret_password"
	hashedPassword, err := security.Hash(password)
	if err != nil {
		t.Errorf("hash failed, %v", err)
	}
	err = security.VerifyPassword(hashedPassword, password)
	if err != nil {
		t.Errorf("password didn't pass the verification, %v", err)
	}
}
