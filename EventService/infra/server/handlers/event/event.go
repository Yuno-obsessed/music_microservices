package event

import (
	"encoding/json"
	"strconv"

	"github.com/Yuno-obsessed/music_microservices/EventService/domain/entity"
	"github.com/Yuno-obsessed/music_microservices/EventService/service/event"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/lerrors"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/gin-gonic/gin"
)

type Event struct {
	event.EventService
	logger.CustomLogger
}

func NewEvent(service event.EventService) Event {
	return Event{
		service,
		logger.NewLogger(),
	}
}

func (e *Event) EventInfo(c *gin.Context) {
	path := c.Param("id")
	newEvent, err := e.GetOne(path)
	if err != nil {
		e.Logger.Error(lerrors.ErrInQuery.Error() + err.Error())
		c.AbortWithStatusJSON(400, err.Error())
	}
	err = json.NewEncoder(c.Writer).Encode(&newEvent)
	if err != nil {
		e.Logger.Error(lerrors.ErrMarshallingJson.Error() + err.Error())
		c.AbortWithStatusJSON(500, err.Error())
	}
	e.Logger.Info("EventInfo handler was called with success")
}

func (e *Event) EventsOfBand(c *gin.Context) {
	path := c.Param("band")
	newEvents, err := e.GetAllOfBand(path)
	if err != nil || newEvents == nil {
		e.Logger.Error(lerrors.ErrInQuery.Error() + err.Error())
		c.AbortWithStatusJSON(400, err.Error())
	}
	err = json.NewEncoder(c.Writer).Encode(&newEvents)
	if err != nil {
		e.Logger.Error(lerrors.ErrMarshallingJson.Error() + err.Error())
		c.AbortWithStatusJSON(500, err.Error())
	}
	e.Logger.Info("EventsOfBand handler was called with success")
}

func (e *Event) EventsOfCity(c *gin.Context) {
	path := c.Param("city")
	newEvents, err := e.GetAllOfCity(path)
	if err != nil || newEvents == nil {
		e.Logger.Error(lerrors.ErrNoRecord.Error() + err.Error())
		c.AbortWithStatusJSON(400, err.Error())
	}
	err = json.NewEncoder(c.Writer).Encode(&newEvents)
	if err != nil {
		e.Logger.Error(lerrors.ErrMarshallingJson.Error() + err.Error())
		c.AbortWithStatusJSON(500, err.Error())
	}
	e.Logger.Info("EventsOfCity was called with success")
}

func (e *Event) EventCreate(c *gin.Context) {
	var newEvent entity.Event
	err := json.NewDecoder(c.Request.Body).Decode(&newEvent)
	if err != nil {
		e.Logger.Error(lerrors.ErrMarshallingJson.Error() + err.Error())
		c.AbortWithStatusJSON(500, err.Error())
	}
	if err = e.EventService.Create(newEvent); err != nil {
		e.Logger.Error(lerrors.ErrInQuery.Error() + err.Error())
		c.AbortWithStatusJSON(400, err.Error())
	}
	e.Logger.Info("EventCreate was called with success")
	c.JSONP(200, "event was created")
}

func (e *Event) EventDelete(c *gin.Context) {
	path := c.Param("id")
	id, err := strconv.Atoi(path)
	if err != nil && id < 0 {
		e.Logger.Error(lerrors.ErrInParam.Error() + err.Error())
		c.AbortWithStatusJSON(400, err.Error())
	}
	err = e.EventService.Delete(path)
	if err != nil {
		e.Logger.Error(lerrors.ErrInQuery.Error() + err.Error())
		c.AbortWithStatusJSON(500, err.Error())
	}
	e.Logger.Info("EventDelete was called with success")
	c.JSONP(200, "event was deleted")
}
