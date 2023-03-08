package upload

import (
	"github.com/Yuno-obsessed/music_microservices/UploadRepository/service/upload/interfaces"
)

type UploadFileService struct {
	Repo interfaces.UploadFile
}

func NewFileUploadService(repo interfaces.UploadFile) *UploadFileService {
	return &UploadFileService{
		repo,
	}
}
