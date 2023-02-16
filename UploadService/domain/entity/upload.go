package entity

type Upload struct {
	UploadId string `json:"upload_id"`
	Name     string `json:"upload_name"`
	Entity   string `json:"upload_entity"`
}
