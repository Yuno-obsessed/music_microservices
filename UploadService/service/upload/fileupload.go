package upload

import (
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/Yuno-obsessed/music_microservices/UploadRepository/service/upload/interfaces"
)

type FileUploadService struct {
	Repo   interfaces.UploadFile
	Logger logger.CustomLogger
}

func NewFileUploadService(repo interfaces.UploadFile) *FileUploadService {
	return &FileUploadService{
		repo,
		logger.NewLogger(),
	}
}
