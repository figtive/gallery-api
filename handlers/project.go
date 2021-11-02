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
	}); err != nil {
		return
	}
	return
}
