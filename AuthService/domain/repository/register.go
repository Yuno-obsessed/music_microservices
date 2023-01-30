package repository

import (
	"auth-service/domain/entity"
	"context"
	"github.com/google/uuid"
)

type RegisterRepository interface {
	SaveRegister(ctx context.Context, register entity.Register) error
	GetRegisterByUuid(ctx context.Context, registerUuid uuid.UUID) (entity.Register, error)
	GetRegisterByEmailAndPassword(ctx context.Context, email string, password string) (entity.Register, error)
	IsRegistered(ctx context.Context, email string, password string) bool // or error
}
