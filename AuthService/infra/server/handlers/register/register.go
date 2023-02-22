package register

import (
	"encoding/json"

	"github.com/Yuno-obsessed/music_microservices/AuthService/domain/dto"
	"github.com/Yuno-obsessed/music_microservices/AuthService/service/register"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Register struct {
	service register.RegisterService
	logger  logger.CustomLogger
}

func (r Register) RegisterUser(c *gin.Context) {
	var role int
	switch c.Param(":role") {
	case "customer":
		role = 0
	case "band":
		role = 1
	case "artist":
		role = 2
	default:
		c.Redirect(308, "/customer")
	}
	var user dto.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		r.logger.Error("error decoding", zap.Error(err))
		c.AbortWithStatusJSON(400, err)
	}

	r.service.SaveRegister(ctx)
}
