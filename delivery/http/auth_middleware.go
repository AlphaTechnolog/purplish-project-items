package http

import (
	"net/http"

	"github.com/alphatechnolog/purplish-items/pkg/helpers"
	"github.com/alphatechnolog/purplish-project-common/auth"
	"github.com/gin-gonic/gin"
)

func getAuthToken() string {
	authToken := helpers.GetEnv("API_GATEWAY_AUTH_TOKEN_B64", "")
	if authToken == "" {
		panic("auth token is required to validate api gateway scopes")
	}
	return authToken
}

func APIGatewayScopeCheck(requiredScopes string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userScopes := c.GetHeader("X-User-Scopes")
		if userScopes == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "X-User-Scopes is required"})
			c.Abort()
			return
		}
		authToken := getAuthToken()
		userScopes, err := auth.ApiGatewayScopeCheck(authToken, userScopes, requiredScopes)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("user-scopes", userScopes)
		c.Next()
	}
}
