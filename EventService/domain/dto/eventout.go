package dto

import "time"

type EventOut struct {
	BandName        string    `json:"band_name"`
	EventCity       string    `json:"event_city"`
	TicketsQuantity int       `json:"tickets_quantity"`
	AveragePrice    int       `json:"average_price"`
	Date            time.Time `json:"event_date"`
}
