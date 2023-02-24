package dto

import "time"

type Mail struct {
	Recipient string    `json:"recipient"`
	Subject   string    `json:"subject"`
	UploadId  int       `json:"upload_id,omitempty"`
	DateSent  time.Time `json:"date_sent"`
}
