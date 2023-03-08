package dto

import "time"

type EventCreate struct {
	BandName         string    `json:"band_name"`
	EventCity        string    `json:"event_city"`
	EventVip         int       `json:"event_vip"`
	EventDefault     int       `json:"event_default"`
	EventScene       int       `json:"event_scene"`
	EventVipCost     int       `json:"event_vip_cost"`
	EventDefaultCost int       `json:"event_default_cost"`
	EventSceneCost   int       `json:"event_scene_cost"`
	Date             time.Time `json:"event_date"`
}
