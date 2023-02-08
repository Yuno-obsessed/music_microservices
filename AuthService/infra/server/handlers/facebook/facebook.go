package facebook

import (
	"auth-service/infra/facebook"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func FacebookLoginHandler(c *gin.Context) {
	facebook.Facebook()
	//gothic.BeginAuthHandler(c.Writer, c.Request)
	provider, err := goth.GetProvider("facebook")
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

func FacebookLoginCallback(c *gin.Context) {
	facebook.Facebook()
	gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.AbortWithStatusJSON(500, fmt.Sprintf("Error calling facebook login, %v", err))
	}
	// sent to endpoint with database
	json.NewEncoder(c.Writer).Encode(gothUser)
}

//func FacebookCallbackHandler(c *gin.Context) {
//	ctx := context.Background()
//	if err != nil {
//		c.AbortWithStatusJSON(500, fmt.Sprintf("Error calling facebook login, %v", err))
//	}
//
//}
