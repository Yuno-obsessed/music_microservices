package facebook

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func FacebookLoginHandler(c *gin.Context) {
	gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.AbortWithStatusJSON(500, fmt.Sprintf("Error calling facebook login, %v", err))
	}
	// sent to endpoint with database
	json.NewEncoder(c.Writer).Encode(gothUser)
}
