package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter() Router {
	return Router{gin.Default()}
}

func (r Router) CatalogGroup() {
	uploadGroup := r.Group("/api/v1/upload-service")
	uploadGroup.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, "healthy")
	})
}

// Function for testing
func (r Router) InitRoutes() {
	r.CatalogGroup()
	log.Fatal(r.Run(":" + os.Getenv("PORT")))
}

func Init() {
	router := NewRouter()
	router.CatalogGroup()
	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}
