package repository

import (
	"github.com/Yuno-obsessed/music_microservices/AuthService/domain/entity"
)

type RegisterRepository interface {
	SaveRegister(register entity.User) error
	GetRegisterById(registerId int) (entity.User, error)
	GetRegisterByEmailAndPassword(email string, password string) (entity.User, error)
}
