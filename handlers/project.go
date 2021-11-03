package handlers

import (
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"
)

func (m *module) ProjectInsert(projectInfo dtos.ProjectInsert, classID string, thumbnailPath string) (id string, err error) {
	var courseworkID string
	if courseworkID, err = Handler.CourseworkInsert(classID); err != nil {
		return
	}
	if id, err = m.db.projectOrmer.Insert(models.Project{
		CourseworkID: courseworkID,
		Name:         projectInfo.Name,
		Active:       projectInfo.Active,
		Description:  projectInfo.Description,
		Field:        projectInfo.Field,
		Thumbnail:    thumbnailPath,
		Team:         projectInfo.Team,
		Metadata:     projectInfo.Metadata,
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
	for _, project := range projectsRaw {
		projects = append(projects, dtos.Project{
			ID:          project.CourseworkID,
			Name:        project.Name,
			Active:      project.Active,
			Description: project.Description,
			Field:       project.Field,
			Thumbnail:   project.Thumbnail,
			CreatedAt:   project.CreatedAt,
			Team:        project.Team,
			Metadata:    project.Metadata,
		})
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
		Name:        projectRaw.Name,
		Active:      projectRaw.Active,
		Description: projectRaw.Description,
		Field:       projectRaw.Field,
		Thumbnail:   projectRaw.Thumbnail,
		CreatedAt:   projectRaw.CreatedAt,
		Team:        projectRaw.Team,
		Metadata:    projectRaw.Metadata,
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
