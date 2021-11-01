package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/handlers"
)

func POSTLogin(c *gin.Context) {
	var err error
	var body dtos.UserLogin
	if err = c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err})
		return
	}

	var claims dtos.GoogleJWTClaim
	if claims, err = handlers.Handler.AuthParseGoogleJWT(body.Token); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err})
		return
	}

	var userInfo dtos.User
	if userInfo, err = handlers.Handler.UserGetOneByEmail(claims.Email); err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err})
		return
	}

	if userInfo.ID == "" {
		userInfo.Email = claims.Email
		userInfo.Name = claims.Name
		if userInfo.ID, err = handlers.Handler.UserInsert(userInfo); err != nil {
			c.JSON(http.StatusInternalServerError, dtos.Response{Error: err})
			return
		}
	}

	var token string
	if token, err = handlers.Handler.AuthGenerateJWT(userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err})
		return
	}

	c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	c.JSON(http.StatusOK, dtos.Response{
		Code: 200,
	})
}
