package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
)

func POSTLogin(c *gin.Context) {
	c.JSON(http.StatusOK, dtos.Response{
		Code: 200,
	})
}
