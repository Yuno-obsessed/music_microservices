package server

import (
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

func (r Router) CustomerGroup() {
	mailGroup := r.Group("api/v1/mail")
	//mailGroup.GET("/:message_type", mailing.MailSuccessfulRegistration)
	mailGroup.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, "healthy")
	})
}

// Function for testing
func (r Router) InitRoutes() {
	r.CustomerGroup()
	log.Fatal(r.Run(":" + os.Getenv("PORT")))
}

func Init() {
	router := NewRouter()
	router.CustomerGroup()
	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}
