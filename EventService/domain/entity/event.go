package entity

import (
	"customer-service/domain/entity"
	"time"
)

type Event struct {
	EventId   string      `json:"event_id"`
	BandName  string      `json:"band_name"`
	EventCity entity.City `json:"event_city"`
	Date      time.Time   `json:"date"`
}
