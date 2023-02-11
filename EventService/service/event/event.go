package event

import (
	"event-service/domain/entity"
	"fmt"
	"github.com/jinzhu/gorm"
)

type EventService struct {
	db *gorm.DB
}

func NewEventService(db *gorm.DB) *EventService {
	return &EventService{db}
}

func (es *EventService) GetOne(id string) (entity.Event, error) {
	var event entity.Event
	err := es.db.Debug().Where("event_id=?", id).Take(&event).Error
	if err != nil {
		return entity.Event{}, fmt.Errorf("such event wasn't found in database, %v", err)
	}
	return event, nil
}

func (es *EventService) GetAllOfBand(band string) ([]entity.Event, error) {
	var events []entity.Event
	err := es.db.Debug().Limit(10).Order("desc").Where("band_name=?", band).Take(&events).Error
	if err != nil {
		return nil, fmt.Errorf("such event wasn't found in database, %v", err)
	}
	return events, nil
}

// add limit and offset
func (es *EventService) GetAll() ([]entity.Event, error) {
	var events []entity.Event
	err := es.db.Debug().Limit(10).Order("desc").Take(&events).Error
	if err != nil {
		return nil, fmt.Errorf("such event wasn't found in database, %v", err)
	}
	return events, nil
}

func (es *EventService) Create(event entity.Event) error {
	err := es.db.Debug().Create(&event).Error
	return err
}

func (es *EventService) Update(event entity.Event) error {
	err := es.db.Debug().Save(&event).Error
	return err
}

func (es *EventService) Delete(id string) error {
	var event entity.Event
	err := es.db.Debug().Where("event_id=?", id).Delete(&event).Error
	return err
}
