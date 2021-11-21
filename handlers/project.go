package handlers

import (
	"math/rand"
	"time"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/configs"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/utils"
)

func (m *module) ProjectInsert(projectInfo dtos.ProjectInsert, classID string, thumbnailPath string) (id string, err error) {
	var courseworkID string
	if courseworkID, err = Handler.CourseworkInsert(classID); err != nil {
		return
	}
	if id, err = m.db.projectOrmer.Insert(models.Project{
		CourseworkID: courseworkID,
		Name:         projectInfo.Name,
		Team:         projectInfo.Team,
		Description:  projectInfo.Description,
		Thumbnail:    thumbnailPath,
		Field:        projectInfo.Field,
		Active:       projectInfo.Active,
		Metadata:     *projectInfo.Metadata,
	}); err != nil {
		return
	}
	return
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
			Thumbnail:   configs.AppConfig.StaticBaseURL + project.Thumbnail,
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
		Field:       projectRaw.Field,
		Active:      projectRaw.Active,
		Metadata:    projectRaw.Metadata,
		CreatedAt:   projectRaw.CreatedAt,
	}
	return
}

func (m *module) ProjectUpdateThumbnail(id string, thumbnailPath string) error {
	var err error
	if err = m.db.projectOrmer.UpdateThumbnail(id, thumbnailPath); err != nil {
		return err
	}
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
			Field:       project.Field,
			Active:      project.Active,
			Metadata:    project.Metadata,
			CreatedAt:   project.CreatedAt,
		}
	}
	return projects, nil
}
