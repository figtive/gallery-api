package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/handlers"
)

func POSTClass(c *gin.Context) {
	var err error

	var classInfo dtos.Class
	if err := c.ShouldBindJSON(&classInfo); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err})
		return
	}

	if classInfo.ID, err = handlers.Handler.ClassInsert(classInfo); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err})
		return
	}

	c.JSON(http.StatusOK, dtos.Response{
		Code: http.StatusOK,
		Data: classInfo,
	})
}
