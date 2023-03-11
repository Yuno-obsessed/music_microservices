package interfaces

import (
	"github.com/Yuno-obsessed/music_microservices/EventService/domain/dto"
)

type EventInterface interface {
	GetOne(id int) (dto.Event, error)
	GetAllOfBand(band string) ([]dto.Event, error)
	GetAllOfCity(city string) ([]dto.Event, error)
	// add limit and offset
	GetAll() ([]dto.Event, error)
	Create(event dto.EventCreate) (int, error)
	Update(id int, event dto.Event) error
	Delete(id int) error
}
