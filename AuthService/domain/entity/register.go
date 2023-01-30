package entity

import "github.com/google/uuid"

type Register struct {
	UUID     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Country  string    `json:"country"`
}
