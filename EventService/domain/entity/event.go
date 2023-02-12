package entity

import (
	"time"
)

type Event struct {
	EventId     string    `json:"event_id"`
	BandName    string    `json:"band_name"`
	EventCityID string    `json:"event_city_id"`
	Date        time.Time `json:"event_date"`
	//format := "2006-01-02"
	//date, _ := time.Parse(format, "2019-07-10")
}
