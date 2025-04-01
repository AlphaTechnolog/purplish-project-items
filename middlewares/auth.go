package middlewares

import (
	"net/http"

	"github.com/alphatechnolog/purplish-project-common/auth"
	"github.com/gin-gonic/gin"
)

// TODO: Extract these to envvars.
const API_GATEWAY_AUTH_TOKEN_B64 = "BTcZcmbaQDMkRt5gtdQ9c/c2mpEc1ZPehUm1KEOU7oE="

// Middleware that will check for required scopes on the
// scopes redirect by api gateway, we should get a header called
// X-User-Scopes which will contain the user scopes
func APIGatewayScopeCheck(requiredScopes string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userScopes := c.GetHeader("X-User-Scopes")
		if userScopes == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "X-User-Scopes required"})
			c.Abort()
			return
		}
		userScopes, err := auth.ApiGatewayScopeCheck(API_GATEWAY_AUTH_TOKEN_B64, userScopes, requiredScopes)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("user-scopes", userScopes)
		c.Next()
	}
}
