package handlers

import (
	"time"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"
)

func (m *module) LeaderboardProject(term time.Time, courseID string) ([]dtos.Project, error) {
	var err error
	var rawProjects []models.Project
	if rawProjects, err = m.db.projectOrmer.GetManyByTermAndCourseIdSortByVotes(term, courseID); err != nil {
		return nil, err
	}
	projects := make([]dtos.Project, len(rawProjects))
	for i, rawProject := range rawProjects {
		projects[i] = dtos.Project{
			ID:          rawProject.CourseworkID,
			CourseID:    rawProject.Coursework.CourseID,
			Name:        rawProject.Name,
			Team:        rawProject.Team,
			Description: rawProject.Description,
			Thumbnail:   rawProject.Thumbnail,
			Field:       rawProject.Field,
			Active:      rawProject.Active,
			Metadata:    rawProject.Metadata,
			CreatedAt:   rawProject.CreatedAt,
		}
	}
	return projects, nil
}
