package repository

import "github.com/Yuno-obsessed/music_microservices/EventService/domain/entity"

type EventRepository interface {
	GetOne(id string) (entity.Event, error)
	GetAllOfBand(band string) ([]entity.Event, error)
	// add limit and offset
	GetAll() ([]entity.Event, error)
	Create(event entity.Event) error
	Update(event entity.Event) error
	Delete(id string) error
}
