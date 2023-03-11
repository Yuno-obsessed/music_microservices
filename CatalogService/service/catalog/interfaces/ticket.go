package interfaces

import (
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/consts"
	"projects/music_microservices/StorageService/domain/dto"
)

type TicketInterface interface {
	GetSumAndAverage(id int) (dto.TicketBrief, error)
	GetEntity(id int) (dto.TicketOut, error)
	Subtruct(id int, ttype consts.TicketType, number int) error
}
