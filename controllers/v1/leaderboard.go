package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/handlers"
)

func GETProjectLeaderboard(c *gin.Context) {
	var err error

	courseID := c.Param("courseId")

	var projects []dtos.Project
	if projects, err = handlers.Handler.LeaderboardProject(time.Now(), courseID); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK, Data: projects})
}

func GETBlogLeaderboard(c *gin.Context) {
	var err error

	courseID := c.Param("courseId")

	var blogs []dtos.Blog
	if blogs, err = handlers.Handler.LeaderboardBlog(time.Now(), courseID); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK, Data: blogs})
}
