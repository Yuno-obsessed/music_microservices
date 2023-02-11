package event

import (
	"encoding/json"
	"event-service/service/event"
	"github.com/gin-gonic/gin"
)

type Event struct {
	event event.EventService
}

func NewEvent(service event.EventService) *Event {
	return &Event{event: service}
}

func (e *Event) EventInfo(c *gin.Context) {
	path := c.Param(":id")
	newEvent, err := e.event.GetOne(path)
	if err != nil {
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