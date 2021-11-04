package handlers

import (
	"strings"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"
)

func (m *module) CourseInsert(courseInfo dtos.Course) (id string, err error) {
	if id, err = m.db.courseOrmer.Insert(models.Course{
		ID:          courseInfo.ID,
		Name:        courseInfo.Name,
		Description: courseInfo.Description,
	}); err != nil {
		return
	}
	return
}

func (m *module) CourseGetOneByID(id string) (courseInfo dtos.Course, err error) {
	var class models.Course
	if class, err = m.db.courseOrmer.GetOneByID(id); err != nil {
		return dtos.Course{}, err
	}
	courseInfo = dtos.Course{
		ID:          strings.ToLower(class.ID),
		Name:        class.Name,
		Description: class.Description,
	}
	return
}
