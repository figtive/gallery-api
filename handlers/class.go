package handlers

import (
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"
)

func (m *module) ClassInsert(classInfo dtos.Class) (id string, err error) {
	if id, err = m.db.classOrmer.Insert(models.Class{
		ID:          classInfo.ID,
		Name:        classInfo.Name,
		Description: classInfo.Description,
	}); err != nil {
		return
	}
	return
}
