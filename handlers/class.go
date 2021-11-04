package handlers

import (
	"strings"

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

func (m *module) ClassGetOneByID(id string) (classInfo dtos.Class, err error) {
	var class models.Class
	if class, err = m.db.classOrmer.GetOneByID(id); err != nil {
		return dtos.Class{}, err
	}
	classInfo = dtos.Class{
		ID:          strings.ToLower(class.ID),
		Name:        class.Name,
		Description: class.Description,
	}
	return
}
