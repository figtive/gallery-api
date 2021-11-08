package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/constants"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/handlers"
)

func POSTVote(c *gin.Context) {
	var err error
	// check is authed
	if !c.GetBool(constants.ContextIsAuthenticatedKey) {
		c.JSON(http.StatusUnauthorized, dtos.Response{Code: http.StatusUnauthorized, Error: "You must be logged in to vote"})
		return
	}
	// get body
	var voteInsert dtos.VoteInsert
	if err = c.ShouldBindJSON(&voteInsert); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Error: err.Error()})
		return
	}
	// get user from db
	var user dtos.User
	if user, err = handlers.Handler.UserGetOneByEmail(c.GetString(constants.ContextUserEmailKey)); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	// check is coursework exist
	var coursework dtos.Coursework
	if coursework, err = handlers.Handler.CourseworkGetOneByID(voteInsert.CourseworkID); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, dtos.Response{Code: http.StatusNotFound, Error: "Coursework not found"})
		} else {
			c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		}
		return
	}
	// get course object
	var course dtos.Course
	if course, err = handlers.Handler.CourseGetOneByID(coursework.CourseID); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	// get all votes for coursework in current term
	var votes []dtos.Vote
	if votes, err = handlers.Handler.VoteGetVotesForCourseworkInCurrentTerm(user.ID, coursework.ID); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	// check if number of votes has reached quota
	if len(votes) >= course.VoteQuota {
		c.JSON(http.StatusForbidden, dtos.Response{Code: http.StatusForbidden, Error: "You have reached the vote quota"})
		return
	}
	// create vote
	var id string
	if id, err = handlers.Handler.VoteInsert(user.ID, voteInsert); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK, Data: id})
}

func GETVoteCount(c *gin.Context) {
	var err error

	id := c.Param("id")
	var count int64
	if count, err = handlers.Handler.VoteCountByCourseworkID(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, dtos.Response{Code: http.StatusNotFound, Error: "Vote not found"})
		} else {
			c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK, Data: count})
}
