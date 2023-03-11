package event

import (
	"github.com/Yuno-obsessed/music_microservices/EventService/domain/dto"
	"github.com/Yuno-obsessed/music_microservices/EventService/domain/entity"
	"github.com/Yuno-obsessed/music_microservices/EventService/domain/mapping"
	"github.com/Yuno-obsessed/music_microservices/EventService/domain/repository"
)

type EventService struct {
	repo repository.EventRepository
}

func NewEventService() EventService {
	return EventService{
		*repository.NewEventRepository(),
	}
}

func (e *EventService) GetOne(id int) (entity.Event, error) {
	event, err := e.repo.GetOne(id)
	if err != nil {
		return entity.Event{}, err
	}
	return event, nil
}

func (e *EventService) CreateEvent(event dto.EventCreate) (int, error) {
	id, err := e.repo.Create(mapping.EventCreateToDto(event))
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (e *EventService) GetAllInCity(city string) ([]entity.Event, error) {
	events, err := e.repo.GetAllOfCity(city)
	if err != nil {
		return nil, err
	}
	return events, nil
}
