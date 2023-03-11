package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"projects/music_microservices/StorageService/infra/server/handlers"
)

type Router struct {
	*gin.Engine
}

func NewRouter() Router {
	return Router{gin.Default()}
}

func (r Router) CatalogGroup() {
	catalog := handlers.NewTicket()
	catalogGroup := r.Group("/api/v1/catalog-service")
	catalogGroup.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, "healthy")
	})
	catalogGroup.GET("/ticket/{id}/info", catalog.GetEntity)
	catalogGroup.GET("/ticket/{id}/brief", catalog.GetSumAndAverage)
	catalogGroup.GET("/ticket/{id}/subtract", catalog.Subtract)
}

// Function for testing
func (r Router) InitRoutes() {
	r.CatalogGroup()
	log.Fatal(r.Run(":" + os.Getenv("CATALOG_PORT")))
}

func Init() {
	router := NewRouter()
	router.CatalogGroup()
	log.Fatal(router.Run(":" + os.Getenv("CATALOG_PORT")))
}
