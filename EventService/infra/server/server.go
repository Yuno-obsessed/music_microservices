package server

import (
	"github.com/Yuno-obsessed/music_microservices/EventService/infra/server/handlers/event"
	"github.com/Yuno-obsessed/music_microservices/EventService/service"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type Router struct {
	*gin.Engine
}

func NewRouter() Router {
	return Router{gin.Default()}
}

func (r Router) EventGroup() {
	eventGroup := r.Group("api/v1/event")
	repos := service.NewRepositories()
	eventRepo := event.NewEvent(repos.Event)
	eventGroup.GET("/:band/get", eventRepo.EventsOfBand)
	eventGroup.GET("/:city/get", eventRepo.EventsOfCity)
	eventGroup.GET("/:id", eventRepo.EventInfo)
	eventGroup.GET("/create", eventRepo.EventCreate)
	eventGroup.GET("/:id/delete", eventRepo.EventDelete)
	eventGroup.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, "healthy")
	})
}

// Function for testing
func (r Router) InitRoutes() {
	r.EventGroup()
	log.Fatal(r.Run(":" + os.Getenv("PORT")))
}

func Init() {
	router := NewRouter()
	router.EventGroup()
	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}
