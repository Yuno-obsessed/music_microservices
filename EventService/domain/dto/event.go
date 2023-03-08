package dto

import "time"

type Event struct {
	BandName  string    `json:"band_name"`
	EventCity string    `json:"event_city"`
	Date      time.Time `json:"event_date"`
	// format := "2006-01-02"
	// date, _ := time.Parse(format, "2019-07-10")
}
