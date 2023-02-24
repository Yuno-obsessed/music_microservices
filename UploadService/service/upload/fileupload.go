package upload

import (
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/Yuno-obsessed/music_microservices/UploadRepository/service/upload/interfaces"
)

type UploadFileService struct {
	Repo   interfaces.UploadFile
	Logger logger.CustomLogger
}

func NewFileUploadService(repo interfaces.UploadFile) *UploadFileService {
	return &UploadFileService{
		repo,
		logger.NewLogger(),
	}
}
