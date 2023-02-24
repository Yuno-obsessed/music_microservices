package interfaces

import (
	"github.com/Yuno-obsessed/music_microservices/MailService/domain/dto"
)

type Mail interface {
	Save(mail dto.Mail) error
	Delete(id int) error
	DeleteAllOfRecipient(email string) error
	GetByRecipient(email string) ([]dto.Mail, error)
}
