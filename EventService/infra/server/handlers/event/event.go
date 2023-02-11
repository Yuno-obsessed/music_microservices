package event

import (
	"encoding/json"
	"github.com/Yuno-obsessed/music_microservices/EventService/service/event"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/gin-gonic/gin"
)

type Event struct {
	event event.EventService
	logger logger.CustomLogger
}

func NewEvent(service event.EventService) Event {
	return Event{
		event: service,
		logger: logger.NewLogger(),
	}
}

func (e *Event) EventInfo(c *gin.Context) {
	path := c.Param(":id")
	newEvent, err := e.event.GetOne(path)
	if err != nil {
		e.logger
		c.AbortWithStatusJSON(400, err.Error())
	}
	err = json.NewEncoder(c.Writer).Encode(&newEvent)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
}

func (e *Event) EventsOfBand(c *gin.Context) {
	path := c.Param(":band")
	newEvents, err := e.event.GetAllOfBand(path)
	if err != nil || newEvents == nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	err = json.NewEncoder(c.Writer).Encode(&newEvents)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
}

func ()