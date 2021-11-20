package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/constants"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/handlers"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/utils"
)

func GoogleOAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		if jwtString == "" {
			c.Set(constants.ContextIsAuthenticatedKey, false)
			c.Set(constants.ContextIsAdminKey, false)
			c.Next()
			return
		}

		if valid, claims, err := utils.ValidateGoogleJWT(jwtString); err == nil {
			if valid {
				if user, err := handlers.Handler.UserGetOneByEmail(claims.Email); err == nil {
					c.Set(constants.ContextIsAuthenticatedKey, true)
					c.Set(constants.ContextGoogleJWTKey, jwtString)
					c.Set(constants.ContextUserEmailKey, user.Email)
					c.Set(constants.ContextIsAdminKey, user.IsAdmin)
				} else {
					log.Println(err)
					c.AbortWithStatus(http.StatusUnauthorized)
				}
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

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !c.GetBool(constants.ContextIsAdminKey) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
