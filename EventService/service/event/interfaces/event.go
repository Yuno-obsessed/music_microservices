package interfaces

import (
	"github.com/Yuno-obsessed/music_microservices/EventService/domain/dto"
)

type EventInterface interface {
	GetOne(id int) (dto.EventOut, error)
	GetAllOfBand(band string) ([]dto.EventOut, error)
	GetAllOfCity(city string) ([]dto.EventOut, error)
	// add limit and offset
	GetAll() ([]dto.EventOut, error)
	Create(event dto.EventCreate) (int, error)
	Update(id int, event dto.Event) error
	Delete(id int) error
}
