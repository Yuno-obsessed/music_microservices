package entity

type Ticket struct {
	EventId         int `json:"event_id"`
	DefaultQuantity int `json:"default_quantity"`
	VipQuantity     int `json:"vip_quantity"`
	SceneQuantity   int `json:"scene_quantity"`
	DefaultCost     int `json:"default_cost"`
	VipCost         int `json:"vip_cost"`
	SceneCost       int `json:"scene_cost"`
}
