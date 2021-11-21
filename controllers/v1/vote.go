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

	// get coursework id
	voteInsert := dtos.VoteInsert{
		CourseworkID: c.Param("coursework_id"),
	}
	// get post body
	var voteVote dtos.VoteVote
	if err = c.ShouldBindJSON(&voteVote); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Error: err.Error()})
		return
	}
	// get user from db
	var user dtos.User
	if user, err = handlers.Handler.UserGetOneByEmail(c.GetString(constants.ContextUserEmailKey)); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	// if user want to vote
	if voteVote.Vote {
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
		// check if user has voted for this coursework
		var hasVoted bool
		if hasVoted, err = handlers.Handler.VoteHasVoted(user.ID, voteInsert.CourseworkID); err != nil {
			c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
			return
		}
		if hasVoted {
			c.JSON(http.StatusForbidden, dtos.Response{Code: http.StatusForbidden, Error: "You have already voted"})
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
		return
	}
	// else, user want to unvote
	if err = handlers.Handler.VoteUnvote(user.ID, voteInsert.CourseworkID); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, dtos.Response{Code: http.StatusNotFound, Error: "user hasn't voted this coursework"})
		} else {
			c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK})
}

func GETVoteCount(c *gin.Context) {
	var err error

	id := c.Param("coursework_id")
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

func GETVoteQuota(c *gin.Context) {
	var err error
	// update this when adding new coursework type
	cwType := []string{"projects", "blogs"}

	var user dtos.User
	if user, err = handlers.Handler.UserGetOneByEmail(c.GetString(constants.ContextUserEmailKey)); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	quotas := make(map[string]map[string]int)
	var rawCourses []dtos.Course
	if rawCourses, err = handlers.Handler.CourseGetAll(); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	courseQuotaMap := make(map[string]int)
	for _, course := range rawCourses {
		quotas[course.ID] = make(map[string]int)
		courseQuotaMap[course.ID] = course.VoteQuota
		for _, cwType := range cwType {
			quotas[course.ID][cwType] = course.VoteQuota
		}
	}

	for _, cw := range cwType {
		var courseworks []dtos.Coursework
		if courseworks, err = handlers.Handler.CourseworkGetVoted(user.ID, cw); err != nil {
			c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
			return
		}
		for _, coursework := range courseworks {
			quotas[coursework.CourseID][cw]--
		}
	}
	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK, Data: quotas})
}

func GETVoteStatus(c *gin.Context) {
	var err error

	courseworkID := c.Param("coursework_id")
	email := c.GetString(constants.ContextUserEmailKey)
	var user dtos.User
	if user, err = handlers.Handler.UserGetOneByEmail(email); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	if _, err = handlers.Handler.CourseworkGetOneByID(courseworkID); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, dtos.Response{Code: http.StatusNotFound, Error: "Coursework not found"})
		} else {
			c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		}
		return
	}
	var hasVoted bool
	if hasVoted, err = handlers.Handler.VoteHasVoted(user.ID, courseworkID); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK, Data: hasVoted})
}

func GETVotedProject(c *gin.Context) {
	var err error

	userEmail := c.GetString(constants.ContextUserEmailKey)
	var user dtos.User
	if user, err = handlers.Handler.UserGetOneByEmail(userEmail); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	var projects []dtos.Project
	if projects, err = handlers.Handler.VoteGetVotedProjects(user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK, Data: projects})
}
