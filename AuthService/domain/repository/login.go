package repository

import (
	"context"

	"github.com/Yuno-obsessed/music_microservices/AuthService/domain/entity"
)

type LoginRepository interface {
	SaveLogin(ctx context.Context, login entity.Login) error
	GetLoginByEmailAndPassword(ctx context.Context, email string, password string) (entity.Login, error)
}
