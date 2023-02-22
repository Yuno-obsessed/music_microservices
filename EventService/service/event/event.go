package event

import (
	"fmt"
	"strconv"

	"github.com/Masterminds/squirrel"
	"github.com/Yuno-obsessed/music_microservices/EventService/domain/entity"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/database"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EventService struct {
	db     *pgxpool.Pool
	logger logger.CustomLogger
}

func NewEventService() EventService {
	return EventService{
		database.DbInit(),
		logger.NewLogger(),
	}
}

func (es *EventService) GetOne(id string) (entity.Event, error) {
	var event entity.Event

	query, args, err := squirrel.Select()
	err := es.db.Table("events").Debug().
		Where("event_id=?", id).Take(&event).Error
	if err != nil {
		return entity.Event{}, fmt.Errorf("such event wasn't found in database, %v", err)
	}

	return event, nil
}

func (es *EventService) GetAllOfBand(band string) ([]entity.Event, error) {
	var events []entity.Event

	err := es.db.Table("events").Debug().
		Limit(10).Order("desc").
		Where("band_name=?", band).
		Take(&events).Error
	if err != nil {
		return nil, fmt.Errorf("such event wasn't found in database, %v", err)
	}

	return events, nil
}

func (es *EventService) GetAllOfCity(city string) ([]entity.Event, error) {
	var id int

	err := es.db.Table("cities").Debug().
		Limit(10).Order("desc").
		Where("city_name=?", city).
		Take(&id).Error
	if err != nil {
		_ = es.db.Table("cities").Debug().
			Last(&id).Error

		newCity := entity.City{
			Id:       strconv.Itoa(id + 1),
			CityName: city,
		}
		err := es.db.Table("cities").Debug().
			Create(&newCity).Error
		if err != nil {
			return nil, fmt.Errorf("error creating a city not yet existing in database, %v", err)
		}
	}

	var events []entity.Event

	err = es.db.Table("events").Debug().
		Limit(10).Order("desc").
		Where("event_city_id=?", id).
		Take(&events).Error

	if err != nil {
		return nil, fmt.Errorf("no events were found in this city, %v", err)
	}

	return events, nil
}

// add limit and offset
func (es *EventService) GetAll() ([]entity.Event, error) {
	var events []entity.Event

	err := es.db.Table("events").Debug().
		Limit(10).Order("desc").
		Take(&events).Error
	if err != nil {
		return nil, fmt.Errorf("such event wasn't found in database, %v", err)
	}

	return events, nil
}

func (es *EventService) Create(event entity.Event) error {
	err := es.db.Table("events").Debug().
		Create(&event).Error

	return err
}

func (es *EventService) Update(event entity.Event) error {
	err := es.db.Table("events").Debug().
		Save(&event).Error

	return err
}

func (es *EventService) Delete(id string) error {
	var event entity.Event

	err := es.db.Table("events").Debug().
		Where("event_id=?", id).
		Delete(&event).Error

	return err
}
