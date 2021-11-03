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
		})
	}
	return
}
