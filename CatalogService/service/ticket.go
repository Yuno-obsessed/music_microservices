package service

import (
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/consts"
	"projects/music_microservices/StorageService/domain/dto"
	"projects/music_microservices/StorageService/domain/repository"
)

type TicketService struct {
	Repo *repository.TicketRepository
}

func NewTicketService() *TicketService {
	return &TicketService{
		repository.NewTicketRepository(),
	}
}

func (t *TicketService) GetSumAndAverage(id int) (dto.TicketBrief, error) {
	return t.Repo.GetSumAndAverage(id)
}

func (t *TicketService) GetEntity(id int) (dto.TicketOut, error) {
	return t.Repo.GetEntity(id)
}

func (t *TicketService) Subtruct(id int, ttype consts.TicketType, number int) error {
	return t.Repo.Subtruct(id, ttype, number)
}
