package server

import (
	"log"
	"os"

	"github.com/Yuno-obsessed/music_microservices/AuthService/infra/server/handlers/auth"
	"github.com/Yuno-obsessed/music_microservices/AuthService/infra/server/handlers/facebook"
	"github.com/Yuno-obsessed/music_microservices/AuthService/infra/server/handlers/google"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Router struct {
	*gin.Engine
}

func NewRouter() Router {
	return Router{gin.Default()}
}

func (r Router) AuthGroup() {
	authGroup := r.Group("/api/v1")
	authGroup.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, "healthy")
	})
	authGroup.POST("/login", auth.LoginHandler)
	authGroup.GET("/login/auth/facebook", facebook.FacebookLoginHandler)
	authGroup.GET("/login/auth/facebook/callback", facebook.FacebookLoginCallback)
	authGroup.GET("/login/auth/google", google.GoogleLoginHandler)
	authGroup.GET("/login/auth/google/callback", google.GoogleLoginCallback)
}

// Function for testing
func (r Router) InitRoutes() {
	godotenv.Load("../.env")
	r.AuthGroup()
	log.Fatal(r.Run(":" + os.Getenv("PORT")))
}

func Init() {
	router := NewRouter()
	router.AuthGroup()
	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}
