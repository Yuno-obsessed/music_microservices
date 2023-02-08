package server

import (
	"auth-service/infra/server/handlers/auth"
	"auth-service/infra/server/handlers/facebook"
	"auth-service/infra/server/handlers/google"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
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
