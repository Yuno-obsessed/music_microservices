package event

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Yuno-obsessed/music_microservices/EventService/domain/dto"
	"github.com/Yuno-obsessed/music_microservices/EventService/domain/mapping"
	"github.com/Yuno-obsessed/music_microservices/EventService/service/event"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/lerrors"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Event struct {
	Service event.EventService
	Logger  logger.CustomLogger
}

func NewEvent(service event.EventService) Event {
	return Event{
		service,
		logger.NewLogger(),
	}
}

func (e *Event) EventInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error("error converting param to id", zap.Error(err))
		c.AbortWithStatusJSON(400, gin.H{"error": err})
	}
	newEvent, err := e.Service.GetOne(id)
	if err != nil {
		e.Logger.Error("error invoking service", zap.Error(err))
		c.AbortWithStatusJSON(400, gin.H{"error": err})
	}
	req, err := http.NewRequest("POST", "http://localhost:8088/api/v1/catalog-service/ticket/"+c.Param("id"), nil)
	if err != nil {
		e.Logger.Error(lerrors.ErrMakingRequest, zap.Error(err))
		c.AbortWithStatusJSON(500, gin.H{"error": err})
	}
	if req.Response.StatusCode != http.StatusOK {
		e.Logger.Error("error sending a request to create a ticket, got not-OK status")
		c.AbortWithStatusJSON(500, gin.H{"error": "response status isn't 200"})
	}
	var ticketInfo dto.TicketInfo
	err = json.NewDecoder(req.Body).Decode(&ticketInfo)
	quantity := (ticketInfo.EventDefault + ticketInfo.EventVip + ticketInfo.EventScene) / 3
	average := (ticketInfo.EventDefaultCost + ticketInfo.EventVipCost + ticketInfo.EventSceneCost) / 3
	eventOut := mapping.EventToEventInfo(newEvent, quantity, average)

	err = json.NewEncoder(c.Writer).Encode(&eventOut)
	if err != nil {
		e.Logger.Error(lerrors.ErrMarshallingJson.Error(), zap.Error(err))
		c.AbortWithStatusJSON(500, gin.H{"error": err})
	}
	e.Logger.Info("EventInfo handler was called with success")
}

func (e *Event) EventsOfBand(c *gin.Context) {
	//path := c.Param("band")
	//newEvents, err := e.GetAllOfBand(path)
	//if err != nil || newEvents == nil {
	//	e.Logger.Error(lerrors.ErrInQuery.Error() + err.Error())
	//	c.AbortWithStatusJSON(400, err.Error())
	//}
	//err = json.NewEncoder(c.Writer).Encode(&newEvents)
	//if err != nil {
	//	e.Logger.Error(lerrors.ErrMarshallingJson.Error() + err.Error())
	//	c.AbortWithStatusJSON(500, err.Error())
	//}
	//e.Logger.Info("EventsOfBand handler was called with success")
}

func (e *Event) EventsOfCity(c *gin.Context) {
	//path := c.Param("city")
	//newEvents, err := e.GetAllOfCity(path)
	//if err != nil || newEvents == nil {
	//	e.Logger.Error(lerrors.ErrNoRecord.Error() + err.Error())
	//	c.AbortWithStatusJSON(400, err.Error())
	//}
	//err = json.NewEncoder(c.Writer).Encode(&newEvents)
	//if err != nil {
	//	e.Logger.Error(lerrors.ErrMarshallingJson.Error() + err.Error())
	//	c.AbortWithStatusJSON(500, err.Error())
	//}
	//e.Logger.Info("EventsOfCity was called with success")
}

func (e *Event) EventCreate(c *gin.Context) {
	var newEvent dto.EventCreate
	err := json.NewDecoder(c.Request.Body).Decode(&newEvent)
	if err != nil {
		e.Logger.Error(lerrors.ErrMarshallingJson.Error(), zap.Error(err))
		c.AbortWithStatusJSON(500, gin.H{"error": err})
	}
	id, err := e.Service.CreateEvent(newEvent)
	if err != nil {
		e.Logger.Error(lerrors.ErrInQuery.Error(), zap.Error(err))
		c.AbortWithStatusJSON(400, gin.H{"error": err})
	}
	payload, err := json.Marshal(mapping.EventCreateToTicketInfo(id, newEvent))
	// requests should be in handlers layer.
	req, err := http.NewRequest("POST", "https://localhost:8088/api/v1/catalog-service/ticket/create", bytes.NewBuffer(payload))
	if err != nil {
		e.Logger.Error(lerrors.ErrMakingRequest, zap.Error(err))
		c.AbortWithStatusJSON(500, gin.H{"error": err})
	}
	if req.Response.StatusCode != http.StatusOK {
		e.Logger.Error("error sending a request to create a ticket, got not-OK status")
		c.AbortWithStatusJSON(500, gin.H{"error": "response status isn't 200"})
	}
	e.Logger.Info("EventCreate was called with success")
	c.JSONP(200, "event was created")
}

func (e *Event) EventDelete(c *gin.Context) {
	//path := c.Param("id")
	//id, err := strconv.Atoi(path)
	//if err != nil && id < 0 {
	//	e.Logger.Error(lerrors.ErrInParam.Error() + err.Error())
	//	c.AbortWithStatusJSON(400, err.Error())
	//}
	//err = e.EventService.Delete(path)
	//if err != nil {
	//	e.Logger.Error(lerrors.ErrInQuery.Error() + err.Error())
	//	c.AbortWithStatusJSON(500, err.Error())
	//}
	//e.Logger.Info("EventDelete was called with success")
	//c.JSONP(200, "event was deleted")
}
