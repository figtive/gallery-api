package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/constants"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/handlers"
	"gorm.io/gorm"
)

func POSTBookmark(c *gin.Context) {
	var err error
	var action dtos.BookmarkAction
	if err = c.ShouldBindJSON(&action); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Error: err.Error()})
		return
	}

	var user dtos.User
	if user, err = handlers.Handler.UserGetOneByEmail(c.GetString(constants.ContextUserEmailKey)); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	bookmark := dtos.Bookmark{
		UserID:       user.ID,
		CourseworkID: c.Param("coursework_id"),
	}

	if _, err = handlers.Handler.CourseworkGetOneByID(bookmark.CourseworkID); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, dtos.Response{Code: http.StatusNotFound, Error: "coursework not found"})
		} else {
			c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		}
		return
	}

	if action.Mark {
		var hasMarked bool
		if hasMarked, err = handlers.Handler.BookmarkHasMarked(bookmark); err != nil {
			c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
			return
		}
		if hasMarked {
			c.JSON(http.StatusForbidden, dtos.Response{Code: http.StatusForbidden, Error: "user has bookmarked this coursework"})
			return
		}
		if _, err = handlers.Handler.BookmarkInsert(bookmark); err != nil {
			c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
			return
		}
	} else {
		if err = handlers.Handler.BookmarkDelete(bookmark); err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, dtos.Response{Code: http.StatusForbidden, Error: "user has not bookmarked this coursework"})
			} else {
				c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
			}
			return
		}
	}
	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK})
}

func GETBookmarkBlogs(c *gin.Context) {
	var err error
	var user dtos.User
	if user, err = handlers.Handler.UserGetOneByEmail(c.GetString(constants.ContextUserEmailKey)); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	var blogs []dtos.Blog
	if blogs, err = handlers.Handler.BookmarkGetManyBlogByUserID(user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK, Data: blogs})
}

func GETBookmarkProjects(c *gin.Context) {
	var err error
	var user dtos.User
	if user, err = handlers.Handler.UserGetOneByEmail(c.GetString(constants.ContextUserEmailKey)); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	var projects []dtos.Project
	if projects, err = handlers.Handler.BookmarkGetManyProjectByUserID(user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK, Data: projects})
}

func GETBookmarkStatus(c *gin.Context) {
	var err error

	courseworkID := c.Param("coursework_id")

	var user dtos.User
	if user, err = handlers.Handler.UserGetOneByEmail(c.GetString(constants.ContextUserEmailKey)); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	if _, err = handlers.Handler.CourseworkGetOneByID(courseworkID); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, dtos.Response{Code: http.StatusNotFound, Error: "coursework not found"})
		} else {
			c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		}
		return
	}

	var hasVoted bool
	if hasVoted, err = handlers.Handler.BookmarkHasMarked(dtos.Bookmark{
		UserID:       user.ID,
		CourseworkID: courseworkID,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK, Data: hasVoted})
}
