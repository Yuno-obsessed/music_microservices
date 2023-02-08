package auth

import (
	"auth-service/domain/entity"
	"auth-service/infra/auth"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// LoginHandler is get request with payload of json{"email", "password"}
func LoginHandler(c *gin.Context) {
	var login entity.Login
	json.NewDecoder(c.Request.Body).Decode(&login)
	jwt, err := auth.NewJWT().GenerateToken(login.Email)
	if err != nil {
		c.AbortWithStatusJSON(500, fmt.Sprintf("error "))
	}
	c.SetCookie("Authorization", jwt,
		int(time.Now().Add(72*time.Hour).Unix()),
		"", "", false, true)
}