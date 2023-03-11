package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/lerrors"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"projects/music_microservices/StorageService/domain/dto"
	"projects/music_microservices/StorageService/service"
)

type Ticket struct {
	Service *service.TicketService
	Logger  logger.CustomLogger
}

func NewTicket() *Ticket {
	return &Ticket{
		service.NewTicketService(),
		logger.NewLogger(),
	}
}

func (t *Ticket) GetSumAndAverage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		t.Logger.Error("err in params", zap.Error(err))
		c.AbortWithStatusJSON(400, gin.H{"error": err})
	}
	dto, err := t.Service.GetSumAndAverage(id)
	if err != nil {
		t.Logger.Error("error:", zap.Error(err))
		c.AbortWithStatusJSON(500, gin.H{"error": err})
	}
	err = json.NewEncoder(c.Writer).Encode(&dto)
	if err != nil {
		t.Logger.Error("error encoding dto:", zap.Error(err))
		c.AbortWithStatusJSON(400, gin.H{"error": err})
	}
	t.Logger.Info("GetSumAndAverage was called with success")
	c.JSONP(200, gin.H{"info": "success"})
}

func (t *Ticket) GetEntity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		t.Logger.Error("err in params", zap.Error(err))
		c.AbortWithStatusJSON(400, gin.H{"error": err})
	}
	dto, err := t.Service.GetEntity(id)
	if err != nil {
		t.Logger.Error("error:", zap.Error(err))
		c.AbortWithStatusJSON(500, gin.H{"error": err})
	}
	err = json.NewEncoder(c.Writer).Encode(&dto)
	if err != nil {
		t.Logger.Error("error encoding dto:", zap.Error(err))
		c.AbortWithStatusJSON(400, gin.H{"error": err})
	}
	t.Logger.Info("GetEntity was called with success")
	c.JSONP(200, gin.H{"info": "success"})
}

func (t *Ticket) Subtract(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		t.Logger.Error("err in params", zap.Error(err))
		c.AbortWithStatusJSON(400, gin.H{"error": err})
	}
	var ticket dto.TicketSubtract
	err = json.NewDecoder(c.Request.Body).Decode(&ticket)
	if err != nil {
		t.Logger.Error(lerrors.ErrMarshallingJson.Error(), zap.Error(err))
		c.AbortWithStatusJSON(400, gin.H{"error": err})
	}
	err = t.Service.Subtruct(id, ticket.TicketType, ticket.TicketType)
	if err != nil {
		t.Logger.Error("error:", zap.Error(err))
		c.AbortWithStatusJSON(500, gin.H{"error": err})
	}
	t.Logger.Info("Subtract was called with success")
	c.JSONP(200, gin.H{"info": "success subtracting "})
}
