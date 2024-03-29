package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

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

	var query dtos.BlogQuery
	if err = c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err.Error()})
		return
	}

	query.CourseID = c.Param("course_id")
	var blogs []dtos.Blog
	if blogs, err = handlers.Handler.BlogGetMany(query.Skip, query.Limit, query.CourseID, query.Title, query.Category, query.Current); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK, Data: blogs})
}

func PUTBlog(c *gin.Context) {
	var err error

	var blogUpdate dtos.BlogUpdate
	if err = c.ShouldBindJSON(&blogUpdate); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err.Error()})
		return
	}

	if err = handlers.Handler.BlogUpdate(blogUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, dtos.Response{Code: http.StatusNotFound, Error: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, dtos.Response{Error: err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK})
}

func DELETEBlog(c *gin.Context) {
	var err error

	blogID := c.Param("blog_id")
	if err = handlers.Handler.BlogDelete(blogID); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, dtos.Response{Code: http.StatusNotFound, Error: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, dtos.Response{Error: err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK})
}
