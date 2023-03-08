package handlers

import (
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"projects/music_microservices/StorageService/service"
)

type Ticket struct {
	Service *service.TicketService
	Logger  logger.CustomLogger
}

func NewTicket() *Ticket {
	return &Ticket{
		service.NewTicketService(),
		logger.NewLogger(),
	}
}
