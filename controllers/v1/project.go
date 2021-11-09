package v1

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/configs"
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
	if projectID, err = handlers.Handler.ProjectInsert(projectInsert, courseInfo.ID, ""); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err.Error()})
		return
	}

	projectInfo := dtos.Project{
		ID:          projectID,
		Name:        projectInsert.Name,
		Active:      projectInsert.Active,
		Description: projectInsert.Description,
		Field:       projectInsert.Field,
		Thumbnail:   "",
	}

	c.JSON(http.StatusOK, dtos.Response{
		Code: http.StatusOK,
		Data: projectInfo,
	})
}

func GETProjects(c *gin.Context) {
	var err error

	var query dtos.Query
	if err = c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err.Error()})
		return
	}
	// TODO: pagination

	var projects []dtos.Project
	if projects, err = handlers.Handler.ProjectGetMany(query.Skip, 0); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.Response{
		Code: http.StatusOK,
		Data: projects,
	})
}

func GETProject(c *gin.Context) {
	var err error

	projectID := c.Param("id")
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

func PUTThumbnail(c *gin.Context) {
	var err error

	var form dtos.ProjectThumbnail
	if err = c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err.Error()})
		return
	}
	var file *multipart.FileHeader
	if file, err = c.FormFile("file"); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err.Error()})
		return
	}

	subdir := fmt.Sprintf("/coursework/project/%s", form.ID)
	filename := fmt.Sprintf("thumbnail%s", filepath.Ext(file.Filename))
	if err = handlers.Handler.ProjectUpdateThumbnail(form.ID, path.Join(subdir, filename)); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err.Error()})
		return
	}
	fullDir := path.Join(configs.AppConfig.StaticBaseDir, subdir)
	if err = os.MkdirAll(fullDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err.Error()})
		return
	}
	if err = c.SaveUploadedFile(file, path.Join(fullDir, filename)); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.Response{
		Code: http.StatusOK,
		Data: form,
	})
}
