package entity

import (
	"time"
)

type Event struct {
	EventId   string    `json:"event_id"`
	BandName  string    `json:"band_name"`
	EventCity City      `json:"event_city"`
	Date      time.Time `json:"date"`
}
