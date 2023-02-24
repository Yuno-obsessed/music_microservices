package upload

import (
	"github.com/Yuno-obsessed/music_microservices/UploadRepository/service/upload/interfaces"
)

type UploadService struct {
	Repo interfaces.Upload
}

func NewUploadService(repo interfaces.Upload) *UploadService {
	return &UploadService{
		Repo: repo,
	}
}
