// I'll have to move all this logic probably to some shared library to use across every micro lmao
package middlewares

import (
	"encoding/base64"
	"net/http"
	"slices"
	"strings"

	"github.com/alphatechnolog/purplish-items/encryption"
	"github.com/gin-gonic/gin"
)

// TODO: Extract these to envvars.
const API_GATEWAY_AUTH_TOKEN_B64 = "BTcZcmbaQDMkRt5gtdQ9c/c2mpEc1ZPehUm1KEOU7oE="

func getApiGatewayAuthToken() []byte {
	contentBytes, err := base64.StdEncoding.DecodeString(API_GATEWAY_AUTH_TOKEN_B64)
	if err != nil {
		panic("Unable to obtain api gateway auth token: " + err.Error())
	}

	return contentBytes
}

// Middleware that will check for required scopes on the
// scopes redirect by api gateway, we should get a header called
// X-User-Scopes which will contain the user scopes
func APIGatewayScopeCheck(requiredScopes []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		encryptedUserScopes := c.GetHeader("X-User-Scopes")
		if encryptedUserScopes == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "X-User-Scopes required"})
			c.Abort()
			return
		}

		userScopes, err := encryption.DecryptAES(getApiGatewayAuthToken(), encryptedUserScopes)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   err.Error(),
				"process": "Obtain user scopes",
			})
			c.Abort()
			return
		}

		c.Set("user-scopes", userScopes)

		splittedScopes := strings.Split(userScopes, " ")
		missingScope := func() *string {
			for _, requiredScope := range requiredScopes {
				if !slices.Contains(splittedScopes, requiredScope) {
					return &requiredScope
				}
			}
			return nil
		}()

		if missingScope != nil {
			// This should be unreachable because the api gateway should've validated this
			// but this failing may indicate one of two things.
			// 1. A possible unknown agent is trying to touch the micro directly and is not passing user scopes correctly which may indicate possible attack attempt.
			// 2. API Gateway is malfunctioning and therefore passing user scopes badly or not validated at all lmfao.
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to perform: " + *missingScope})
			c.Abort()
			return
		}

		c.Next()
	}
}
