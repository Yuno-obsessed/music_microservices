package entity

type Upload struct {
	UploadId string `json:"upload_id"`
	UserId   int    `json:"user_id"`
	Name     string `json:"upload_name"`
	Uentity  string `json:"upload_entity"`
}
