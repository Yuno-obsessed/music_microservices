package entity

import "github.com/Yuno-obsessed/music_microservices/ProjectLibrary/mail"

type Mail struct {
	Id      int              `json:"mail_id"`
	Email   string           `json:"email"`
	Subject string           `json:"subject"`
	Type    mail.MessageType `json:"type"`
}
