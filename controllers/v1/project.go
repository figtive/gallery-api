package v1

import (
	"fmt"
	"mime/multipart"
	"net/http"
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
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err})
		return
	}

	var classInfo dtos.Class
	if classInfo, err = handlers.Handler.ClassGetOneByID(projectInsert.ClassID); err != nil {
		c.JSON(http.StatusNotFound, dtos.Response{Error: "class not found"})
		return
	}

	var projectID string
	if projectID, err = handlers.Handler.ProjectInsert(projectInsert, classInfo.ID, ""); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err})
		return
	}

	projectInfo := dtos.Project{
		ID:          projectID,
		Name:        projectInsert.Name,
		Active:      projectInsert.Active,
		Description: projectInsert.Description,
		Field:       projectInsert.Field,
		// TODO
		Thumbnail: "",
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
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err})
		return
	}
	// TODO: pagination

	var projects []dtos.Project
	if projects, err = handlers.Handler.ProjectGetMany(query.Skip, 0); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err})
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
		c.JSON(http.StatusNotFound, dtos.Response{Code: http.StatusNotFound, Error: "project not found"})
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
		c.JSON(http.StatusBadRequest, dtos.Response{Error: "error parsing form data"})
		return
	}
	var file *multipart.FileHeader
	if file, err = c.FormFile("file"); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Error: "error reading header"})
		return
	}

	subdir := fmt.Sprintf("/coursework/project/%s/thumbnail.png", form.ID)
	if err = handlers.Handler.ProjectUpdateThumbnail(form.ID, subdir); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response{Error: err})
		return
	}
	filename := fmt.Sprintf("thumbnail%s", filepath.Ext(file.Filename))
	c.SaveUploadedFile(file, path.Join(configs.AppConfig.StaticBaseDir, subdir, filename))

	c.JSON(http.StatusOK, dtos.Response{
		Code: http.StatusOK,
		Data: form,
	})
}
