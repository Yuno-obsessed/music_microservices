package dto

import "github.com/Yuno-obsessed/music_microservices/ProjectLibrary/consts"

type TicketSubtract struct {
	TicketType consts.TicketType `json:"ticket_type"`
	Number     int               `json:"number"`
}
