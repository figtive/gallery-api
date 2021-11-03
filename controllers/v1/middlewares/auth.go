package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/constants"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/utils"
)

func GoogleOAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		if jwtString == "" {
			c.Set(constants.ContextIsAuthenticatedKey, false)
			c.Next()
			return
		}

		if valid, claims, err := utils.ValidateGoogleJWT(jwtString); err == nil {
			if valid {
				c.Set(constants.ContextIsAuthenticatedKey, true)
				c.Set(constants.ContextUserEmailKey, claims.Email)
				c.Set(constants.ContextGoogleJWTKey, jwtString)
			} else {
				c.Set(constants.ContextIsAuthenticatedKey, false)
			}
			c.Next()
		} else {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func AuthOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !c.GetBool(constants.ContextIsAuthenticatedKey) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
