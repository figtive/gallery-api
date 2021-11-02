package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/handlers"
)

func POSTBlog(c *gin.Context) {
	var err error

	var blogInsert dtos.BlogInsert
	if err = c.ShouldBindJSON(&blogInsert); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err})
		return
	}

	var classInfo dtos.Class
	if classInfo, err = handlers.Handler.ClassGetOneByID(blogInsert.ClassID); err != nil {
		c.JSON(http.StatusNotFound, dtos.Response{Error: "class not found"})
		return
	}

	// TODO: handle user relation here

	var blogID string
	if blogID, err = handlers.Handler.BlogInsert(blogInsert, classInfo.ID); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err})
		return
	}

	c.JSON(http.StatusOK, dtos.Response{
		Code: http.StatusOK,
		Data: dtos.Blog{
			ID:       blogID,
			Title:    blogInsert.Title,
			Link:     blogInsert.Link,
			Category: blogInsert.Category,
		},
	})
}
