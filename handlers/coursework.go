package handlers

import "gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"

func (m *module) CourseworkInsert(courseID string) (id string, err error) {
	if id, err = m.db.courseworkOrmer.Insert(models.Coursework{
		CourseID: courseID,
	}); err != nil {
		return
	}
	return
}
