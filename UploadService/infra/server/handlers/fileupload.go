package handlers

import (
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/Yuno-obsessed/music_microservices/UploadService/service/upload"
	"github.com/gin-gonic/gin"
)

type FileUpload struct {
	Service upload.FileUploadService
	Logger  logger.CustomLogger
}

func NewFileUpload() FileUpload {
	return FileUpload{
		upload.NewFileUploadService(),
		logger.NewLogger(),
	}
}

func (fu FileUpload) CreateFile(c *gin.Context) {
	// take bucketName from headers or?

	// take file from multipart form?
}
