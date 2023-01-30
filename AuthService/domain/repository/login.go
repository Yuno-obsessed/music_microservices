package repository

import (
	"auth-service/domain/entity"
	"context"
)

type LoginRepository interface {
	SaveLogin(ctx context.Context, login entity.Login) error
	GetLoginByEmailAndPassword(ctx context.Context, email string, password string) (entity.Login, error)
}
