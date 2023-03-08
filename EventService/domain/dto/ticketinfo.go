package dto

type TicketInfo struct {
	EventId          int `json:"event_id"`
	EventDefault     int `json:"event_default"`
	EventVip         int `json:"event_vip"`
	EventScene       int `json:"event_scene"`
	EventDefaultCost int `json:"event_default_cost"`
	EventVipCost     int `json:"event_vip_cost"`
	EventSceneCost   int `json:"event_scene_cost"`
}
