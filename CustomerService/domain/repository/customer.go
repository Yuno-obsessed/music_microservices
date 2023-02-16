package repository

import "github.com/Yuno-obsessed/music_microservices/CustomerService/domain/entity"

type Customer interface {
	GetById(id string) (entity.Customer, error)
	GetByUsername(username string) (entity.Customer, error)
	GetByEmailAndPassword(email string, password string) (entity.Customer, error)
	// add limit and offset
	GetAll() ([]entity.Customer, error)
	Create(event entity.Customer) error
	Update(event entity.Customer) error
	Delete(id string) error
}
