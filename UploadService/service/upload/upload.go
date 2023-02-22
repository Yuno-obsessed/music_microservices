package upload

import (
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/Yuno-obsessed/music_microservices/UploadRepository/service/upload/interfaces"
)

type UploadService struct {
	Repo   interfaces.Upload
	Logger logger.CustomLogger
}

func NewUploadService(repo interfaces.Upload) *UploadService {
	return &UploadService{
		Repo:   repo,
		Logger: logger.NewLogger(),
	}
}
