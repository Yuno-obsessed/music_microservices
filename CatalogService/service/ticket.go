package service

import "projects/music_microservices/StorageService/domain/repository"

type TicketService struct {
	*repository.TicketRepository
}

func NewTicketService() *TicketService {
	return &TicketService{
		repository.NewTicketRepository(),
	}
}
