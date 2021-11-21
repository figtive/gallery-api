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
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err.Error()})
		return
	}

	var courseInfo dtos.Course
	if courseInfo, err = handlers.Handler.CourseGetOneByID(blogInsert.CourseID); err != nil {
		c.JSON(http.StatusNotFound, dtos.Response{Error: err.Error()})
		return
	}

	// TODO: handle user relation here

	var blogID string
	if blogID, err = handlers.Handler.BlogInsert(blogInsert, courseInfo.ID); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err.Error()})
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

func GETBlog(c *gin.Context) {
	var err error

	blogID := c.Param("coursework_id")
	var blog dtos.Blog
	if blog, err = handlers.Handler.BlogGetOne(blogID); err != nil {
		c.JSON(http.StatusNotFound, dtos.Response{Code: http.StatusNotFound, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.Response{
		Code: http.StatusOK,
		Data: blog,
	})
}

func GETBlogs(c *gin.Context) {
	var err error

	var query dtos.Query
	if err = c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err.Error()})
		return
	}

	// TODO: pagination
	var blogs []dtos.Blog
	if blogs, err = handlers.Handler.BlogGetMany(query.Limit, 0); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK, Data: blogs})
}

func GETBlogsInCurrentTermAndCourse(c *gin.Context) {
	var err error

	var query dtos.Query
	if err = c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Error: err.Error()})
		return
	}

	courseID := c.Param("course_id")
	var blogs []dtos.Blog
	if blogs, err = handlers.Handler.BlogGetManyByCourseIDInCurrentTerm(courseID, query.Current); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK, Data: blogs})
}
