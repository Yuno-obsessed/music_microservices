package server

import (
	"log"
	"os"

	"github.com/Yuno-obsessed/music_microservices/UploadService/infra/server/handlers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter() Router {
	return Router{gin.Default()}
}

func (r Router) UploadGroup() {
	upload := handlers.NewUpload()
	uploadGroup := r.Group("/api/v1/upload-service")
	uploadGroup.POST("/new", upload.GetUploadByName)
	uploadGroup.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, "healthy")
	})
}

// Function for testing
func (r Router) InitRoutes() {
	r.UploadGroup()
	log.Fatal(r.Run(":" + os.Getenv("PORT")))
}

func Init() {
	router := NewRouter()
	router.UploadGroup()
	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}
