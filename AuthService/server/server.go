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

func (r Router) AuthGroup() {
	auth := r.Group("/api/v1")
	auth.POST("/login", auth.LoginHandler)
}

func (r Router) RoutesInit() {
	router := NewRouter()
	router.AuthGroup()
	log.Fatal(router.Run(os.Getenv("PORT")))
}
