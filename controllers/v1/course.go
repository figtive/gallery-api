package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/handlers"
)

func POSTCourse(c *gin.Context) {
	var err error

	var courseInfo dtos.Course
	if err = c.ShouldBindJSON(&courseInfo); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err.Error()})
		return
	}

	if courseInfo.ID, err = handlers.Handler.CourseInsert(courseInfo); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.Response{
		Code: http.StatusOK,
		Data: courseInfo,
	})
}

func GETCourses(c *gin.Context) {
	var err error

	var courses []dtos.Course
	if courses, err = handlers.Handler.CourseGetAll(); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK, Data: courses})
}
