package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/handlers"
)

func POSTProject(c *gin.Context) {
	var err error

	var projectInsert dtos.ProjectInsert
	if err = c.ShouldBindJSON(&projectInsert); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err.Error()})
		return
	}

	var courseInfo dtos.Course
	if courseInfo, err = handlers.Handler.CourseGetOneByID(projectInsert.CourseID); err != nil {
		c.JSON(http.StatusNotFound, dtos.Response{Error: err.Error()})
		return
	}

	var projectID string
	if projectID, err = handlers.Handler.ProjectInsert(projectInsert, courseInfo.ID); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err.Error()})
		return
	}

	projectInfo := dtos.Project{
		ID:          projectID,
		Name:        projectInsert.Name,
		Description: projectInsert.Description,
		Thumbnail:   make([]string, 0),
		Field:       projectInsert.Field,
		Active:      projectInsert.Active,
	}

	c.JSON(http.StatusOK, dtos.Response{
		Code: http.StatusOK,
		Data: projectInfo,
	})
}

func GETProject(c *gin.Context) {
	var err error

	projectID := c.Param("coursework_id")
	var projectInfo dtos.Project
	if projectInfo, err = handlers.Handler.ProjectGetOne(projectID); err != nil {
		c.JSON(http.StatusNotFound, dtos.Response{Code: http.StatusNotFound, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.Response{
		Code: http.StatusOK,
		Data: projectInfo,
	})
}

func GETProjects(c *gin.Context) {
	var err error

	var query dtos.ProjectQuery
	if err = c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Error: err.Error()})
		return
	}

	query.CourseID = c.Param("course_id")
	var projects []dtos.Project
	if projects, err = handlers.Handler.ProjectGetMany(query.Skip, query.Limit, query.CourseID, query.Name, query.Field, query.Current); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK, Data: projects})
}

func PUTThumbnail(c *gin.Context) {
	var err error

	var upload dtos.ProjectThumbnailUpload
	if err = c.ShouldBind(&upload); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err.Error()})
		return
	}

	if err = handlers.Handler.ProjectInsertThumbnail(upload.ID, upload.File); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK})
}

func DELETEThumbnail(c *gin.Context) {
	var err error

	var thumbnailDelete dtos.ProjectThumbnailDelete
	if err = c.ShouldBindJSON(&thumbnailDelete); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err.Error()})
		return
	}

	if err = handlers.Handler.ProjectDeleteThumbnail(thumbnailDelete.ID, thumbnailDelete.Thumbnail); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK})
}

func PUTProject(c *gin.Context) {
	var err error

	var projectInfo dtos.ProjectUpdate
	if err = c.ShouldBindJSON(&projectInfo); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err.Error()})
		return
	}

	if err = handlers.Handler.ProjectUpdate(projectInfo); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, dtos.Response{Code: http.StatusNotFound, Error: err.Error()})
			log.Println("here")
		} else {
			c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK})
}

func DELETEProject(c *gin.Context) {
	var err error

	projectID := c.Param("project_id")
	if err = handlers.Handler.ProjectDelete(projectID); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, dtos.Response{Code: http.StatusNotFound, Error: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, dtos.Response{Code: http.StatusInternalServerError, Error: err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK})
}
