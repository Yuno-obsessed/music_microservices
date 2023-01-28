package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kataras/jwt"
	"os"
	"time"
)

type JWT struct {
	Secret []byte
}

func NewJWT() *JWT {
	return &JWT{Secret: []byte(os.Getenv("JWT_SECRET"))}
}

func (j *JWT) GenerateToken(username string) (string, error) {
	claims := jwt.Claims{
		Subject: username,
		Expiry:  time.Now().Add(72 * time.Hour).Unix(),
	}
	token, err := jwt.Sign(jwt.HS256, j.Secret, claims)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func (j *JWT) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Request.Cookie("Authorization")
		tokenString := token.String()[14:]
		if err != nil {
			c.AbortWithStatusJSON(500, fmt.Sprintf("Error getting jwt from cookie, %v", err))
		}
		fmt.Println(tokenString)
		verified, err := jwt.Verify(jwt.HS256, j.Secret, []byte(tokenString))
		if err != nil {
			c.AbortWithStatusJSON(500, fmt.Sprintf("invalid token, %v", err))
			return
		}
		if time.Now().Unix() > verified.StandardClaims.Expiry {
			c.AbortWithStatusJSON(500, fmt.Sprintf("Token expired"))
			return
		}
	}
}
