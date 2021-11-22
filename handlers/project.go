package handlers

import (
	"fmt"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/lib/pq"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/constants"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/utils"
)

func (m *module) ProjectInsert(projectInfo dtos.ProjectInsert, classID string) (string, error) {
	var err error
	var courseworkID string
	if courseworkID, err = Handler.CourseworkInsert(classID, constants.CourseworkTypeProject); err != nil {
		return "", err
	}
	var id string
	if id, err = m.db.projectOrmer.Insert(models.Project{
		CourseworkID: courseworkID,
		Name:         projectInfo.Name,
		Team:         projectInfo.Team,
		Description:  projectInfo.Description,
		Thumbnail:    make(pq.StringArray, 0),
		Link:         projectInfo.Link,
		Video:        projectInfo.Video,
		Field:        projectInfo.Field,
		Active:       projectInfo.Active,
		Metadata:     *projectInfo.Metadata,
	}); err != nil {
		return "", err
	}
	return id, nil
}

func (m *module) ProjectGetMany(skip int, limit int) (projects []dtos.Project, err error) {
	var projectsRaw []models.Project
	if projectsRaw, err = m.db.projectOrmer.GetMany(skip, limit); err != nil {
		return
	}
	projects = make([]dtos.Project, len(projectsRaw))
	for i, j := range rand.Perm(len(projectsRaw)) {
		project := projectsRaw[j]
		projects[i] = dtos.Project{
			ID:          project.CourseworkID,
			CourseID:    project.Coursework.CourseID,
			Name:        project.Name,
			Team:        project.Team,
			Description: project.Description,
			Thumbnail:   project.Thumbnail,
			Link:        project.Link,
			Video:       project.Video,
			Field:       project.Field,
			Active:      project.Active,
			Metadata:    project.Metadata,
			CreatedAt:   project.CreatedAt,
		}
	}
	return
}

func (m *module) ProjectGetOne(id string) (project dtos.Project, err error) {
	var projectRaw models.Project
	if projectRaw, err = m.db.projectOrmer.GetOneByCourseworkID(id); err != nil {
		return
	}
	project = dtos.Project{
		ID:          projectRaw.CourseworkID,
		CourseID:    projectRaw.Coursework.CourseID,
		Name:        projectRaw.Name,
		Team:        projectRaw.Team,
		Description: projectRaw.Description,
		Thumbnail:   projectRaw.Thumbnail,
		Link:        projectRaw.Link,
		Video:       projectRaw.Video,
		Field:       projectRaw.Field,
		Active:      projectRaw.Active,
		Metadata:    projectRaw.Metadata,
		CreatedAt:   projectRaw.CreatedAt,
	}
	return
}

func (m *module) ProjectInsertThumbnail(id string, header *multipart.FileHeader) error {
	var err error
	dir := fmt.Sprintf("/coursework/project/%s/thumbnail-%d%s", id, time.Now().UnixNano(), filepath.Ext(header.Filename))

	if err = utils.SaveMedia(header, dir); err != nil {
		return err
	}

	var project models.Project
	if project, err = m.db.projectOrmer.GetOneByCourseworkID(id); err != nil {
		_ = utils.DeleteMedia(dir)
		return err
	}
	project.Thumbnail = append(project.Thumbnail, dir)
	if err = m.db.projectOrmer.Update(project); err != nil {
		_ = utils.DeleteMedia(dir)
		return err
	}
	return nil
}

func (m *module) ProjectDeleteThumbnail(id string, thumbnailPath string) error {
	var err error
	var project models.Project
	if project, err = m.db.projectOrmer.GetOneByCourseworkID(id); err != nil {
		return err
	}

	var found bool
	for i, thumbnail := range project.Thumbnail {
		if thumbnail == thumbnailPath {
			project.Thumbnail = append(project.Thumbnail[:i], project.Thumbnail[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("thumbnail not found")
	}

	if err = m.db.projectOrmer.Update(project); err != nil {
		return err
	}

	_ = utils.DeleteMedia(thumbnailPath)
	return nil
}

func (m *module) ProjectGetManyByCourseID(courseID string, currentOnly bool) ([]dtos.Project, error) {
	var err error
	var startTime, endTime time.Time
	if currentOnly {
		startTime = utils.TimeToTermTime(time.Now())
		endTime = utils.NextTermTime(time.Now())
	} else {
		startTime = time.Unix(-2208988800, 0)
		endTime = startTime.Add(1<<63 - 1)
	}
	var projectsRaw []models.Project
	if projectsRaw, err = m.db.projectOrmer.GetManyByCourseIDAndTerm(courseID, startTime, endTime); err != nil {
		return nil, err
	}
	projects := make([]dtos.Project, len(projectsRaw))
	for i, j := range rand.Perm(len(projectsRaw)) {
		project := projectsRaw[j]
		projects[i] = dtos.Project{
			ID:          project.CourseworkID,
			CourseID:    project.Coursework.CourseID,
			Name:        project.Name,
			Team:        project.Team,
			Description: project.Description,
			Thumbnail:   project.Thumbnail,
			Link:        project.Link,
			Video:       project.Video,
			Field:       project.Field,
			Active:      project.Active,
			Metadata:    project.Metadata,
			CreatedAt:   project.CreatedAt,
		}
	}
	return projects, nil
}
