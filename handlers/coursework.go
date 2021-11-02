package handlers

import "gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"

func (m *module) CourseworkInsert(classID string) (id string, err error) {
	if id, err = m.db.courseworkOrmer.Insert(models.Coursework{
		ClassID: classID,
	}); err != nil {
		return
	}
	return
}
