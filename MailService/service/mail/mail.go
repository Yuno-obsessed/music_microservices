package mail

import "github.com/Yuno-obsessed/music_microservices/MailService/domain/repository"

// create a service

type MailService struct {
	*repository.MailRepository
}

func NewMailService() *MailService {
	return &MailService{
		repository.NewMailRepository(),
	}
}

func (m *MailService) SendMail()
