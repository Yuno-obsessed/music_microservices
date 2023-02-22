package handlers

import (
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/Yuno-obsessed/music_microservices/UploadService/service/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct {
	Service upload.UploadService
	Logger  logger.CustomLogger
}

func NewUpload() Upload {
	return Upload{
		upload.NewUploadService(),
		logger.NewLogger(),
	}
}

func (u Upload) GetUploadByName(c *gin.Context) {
}
