package handlers

import (
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"
)

func (m *module) CourseworkInsert(courseID string) (id string, err error) {
	if id, err = m.db.courseworkOrmer.Insert(models.Coursework{
		CourseID: courseID,
	}); err != nil {
		return
	}
	return
}

func (m *module) CourseworkGetOneByID(id string) (dtos.Coursework, error) {
	var err error
	var courseworkRaw models.Coursework
	if courseworkRaw, err = m.db.courseworkOrmer.GetOneByID(id); err != nil {
		return dtos.Coursework{}, err
	}
	return dtos.Coursework{
		ID:        courseworkRaw.ID,
		CourseID:  courseworkRaw.CourseID,
		CreatedAt: courseworkRaw.CreatedAt,
	}, nil
}

func (m *module) CourseworkGetVoted(userID, cwType string) ([]dtos.Coursework, error) {
	var err error

	var courseworksRaw []models.Coursework
	if courseworksRaw, err = m.db.courseworkOrmer.GetManyByUserIDAndIsVotedJoinCourseworkType(userID, cwType); err != nil {
		return nil, err
	}
	courseworks := make([]dtos.Coursework, len(courseworksRaw))
	for i, coursework := range courseworksRaw {
		courseworks[i] = dtos.Coursework{
			ID:        coursework.ID,
			CourseID:  coursework.CourseID,
			CreatedAt: coursework.CreatedAt,
		}
	}
	return courseworks, nil
}
