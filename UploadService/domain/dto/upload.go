package dto

type UploadDto struct {
	UserId  int    `json:"user_id"`
	Name    string `json:"upload_name"`
	Uentity string `json:"upload_entity"`
}
