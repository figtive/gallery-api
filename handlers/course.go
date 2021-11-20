package handlers

import (
	"strings"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"
)

func (m *module) CourseInsert(courseInfo dtos.Course) (string, error) {
	var err error
	var id string
	if id, err = m.db.courseOrmer.Insert(models.Course{
		ID:          courseInfo.ID,
		Name:        courseInfo.Name,
		Description: courseInfo.Description,
		VoteQuota:   courseInfo.VoteQuota,
	}); err != nil {
		return "", err
	}
	return id, nil
}

func (m *module) CourseGetOneByID(id string) (dtos.Course, error) {
	var err error
	var courseInfo dtos.Course

	var course models.Course
	if course, err = m.db.courseOrmer.GetOneByID(id); err != nil {
		return courseInfo, err
	}
	courseInfo = dtos.Course{
		ID:          strings.ToLower(course.ID),
		Name:        course.Name,
		Description: course.Description,
		VoteQuota:   course.VoteQuota,
	}
	return courseInfo, nil
}

func (m *module) CourseGetAll() ([]dtos.Course, error) {
	var err error
	var rawCourses []models.Course
	if rawCourses, err = m.db.courseOrmer.GetAll(); err != nil {
		return nil, err
	}
	courses := make([]dtos.Course, len(rawCourses))
	for i, rawCourse := range rawCourses {
		courses[i] = dtos.Course{
			ID:          strings.ToLower(rawCourse.ID),
			Name:        rawCourse.Name,
			Description: rawCourse.Description,
			VoteQuota:   rawCourse.VoteQuota,
		}
	}
	return courses, nil
}
