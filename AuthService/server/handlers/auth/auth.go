package auth

import (
	"auth-service/entity"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var login entity.Login
	json.NewDecoder(c.Request.Body).Decode(&login)
	c.SetCookie("")
}
