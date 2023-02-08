package google

import (
	"auth-service/infra/google"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func GoogleLoginHandler(c *gin.Context) {
	google.Google()
	provider, err := goth.GetProvider("google")
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid provider"})
		return
	}

	// Begin the OAuth flow
	authUrl, err := provider.BeginAuth("")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to begin auth"})
		return
	}

	// Redirect the user to the auth URL
	redirectUrl, err := authUrl.GetAuthURL()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch auth url"})
		return
	}
	c.Redirect(307, redirectUrl)
}

func GoogleLoginCallback(c *gin.Context) {
	google.Google()
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	fmt.Println(user)
	if err != nil {
		c.AbortWithStatusJSON(500, err)
	}
	//json.NewDecoder(c.Request.Body).Decode(data)
	//fmt.Println(data)
}
