package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/constants"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/handlers"
)

func POSTLogin(c *gin.Context) {
	var err error

	var claims dtos.GoogleJWTClaim
	if claims, err = handlers.Handler.AuthParseGoogleJWT(c.GetString(constants.ContextGoogleJWTKey)); err != nil {
		c.JSON(http.StatusUnauthorized, dtos.Response{Error: err.Error()})
		return
	}

	var userInfo dtos.User
	if userInfo, err = handlers.Handler.UserGetOneByEmail(claims.Email); err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err.Error()})
		return
	}

	if userInfo.ID == "" {
		userInfo.Email = claims.Email
		userInfo.Name = claims.Name
		if userInfo.ID, err = handlers.Handler.UserInsert(userInfo); err != nil {
			c.JSON(http.StatusInternalServerError, dtos.Response{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, dtos.Response{
		Code: 200,
	})
}
