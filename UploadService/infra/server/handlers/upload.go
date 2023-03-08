package handlers

import (
	"strconv"
	"strings"

	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/consts"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/Yuno-obsessed/music_microservices/UploadRepository/domain/dto"
	"github.com/Yuno-obsessed/music_microservices/UploadRepository/domain/repository"
	"github.com/Yuno-obsessed/music_microservices/UploadRepository/service/upload"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Upload struct {
	Repository *upload.UploadService
	Upload     *upload.UploadFileService
	Logger     logger.CustomLogger
}

func NewUpload() *Upload {
	return &Upload{
		upload.NewUploadService(
			repository.NewUploadRepository()),
		upload.NewFileUploadService(
			repository.NewFileUploadRepository()),
		logger.NewLogger(),
	}
}

func (u *Upload) SaveUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		u.Logger.Error("error getting file from multipart", zap.Error(err))
		c.AbortWithStatusJSON(400, gin.H{"error": err})
	}
	// 	c.SetCookie("mycookie", user.ID+":"+user.Username+":"+entity, 3600, "/", "", false, true)
	// get username and id from cookies
	cookie, err := c.Cookie("upload")
	if err != nil {
		u.Logger.Error("error processing cookie", zap.Error(err))
		c.AbortWithStatusJSON(500, gin.H{"error processing cookie": err})
	}
	cookieValues := strings.Split(cookie, ":")
	id, err := strconv.Atoi(cookieValues[0])
	if err != nil {
		u.Logger.Error("can't convert string to int", zap.Error(err))
		c.AbortWithStatusJSON(400, gin.H{"error converting string": err})
	}
	newUpload := dto.UploadDto{
		UserId:  id,
		Name:    cookieValues[1],
		Uentity: cookieValues[2],
	}
	_, err = u.Upload.Repo.UploadFile(file, "somename", consts.CustomerBucket)
	if err != nil {
		u.Logger.Error("error uploading file", zap.Error(err))
		c.AbortWithStatusJSON(500, gin.H{"error uploading file": err})
	}
	err = u.Repository.Repo.SaveUpload(newUpload)
	if err != nil {
		u.Logger.Error("error saving upload ", zap.Error(err))
		c.AbortWithStatusJSON(500, gin.H{"error in persistence layer": err})
	}
	c.JSON(200, gin.H{"response": "success saving upload"})
}

// Add get upload method
// How do I make it callable from all the services?
func (u *Upload) GetUploadByName(c *gin.Context) {
}
