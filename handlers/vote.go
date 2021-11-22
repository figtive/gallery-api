package handlers

import (
	"time"

	"gorm.io/gorm"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/utils"
)

func (m *module) VoteInsert(userID string, voteInfo dtos.VoteInsert) (string, error) {
	var err error
	var id string
	if id, err = m.db.voteOrmer.Insert(models.Vote{
		UserID:       userID,
		CourseworkID: voteInfo.CourseworkID,
	}); err != nil {
		return "", err
	}
	return id, nil
}

func (m *module) VoteGetVotesForCourseworkInCurrentTerm(userID, courseworkID string) ([]dtos.Vote, error) {
	var err error
	termDate := utils.TimeToTermTime(time.Now())
	var votesRaw []models.Vote
	if votesRaw, err = m.db.voteOrmer.GetManyByUserIDCourseworkIDAndCreatedAt(userID, courseworkID, termDate); err != nil {
		return nil, err
	}
	votes := make([]dtos.Vote, len(votesRaw))
	for i, vote := range votesRaw {
		votes[i] = dtos.Vote{
			ID:           vote.ID,
			UserID:       vote.UserID,
			CourseworkID: vote.CourseworkID,
		}
	}
	return votes, nil
}

func (m *module) VoteCountByCourseworkID(courseworkID string) (int64, error) {
	var err error
	var count int64
	if count, err = m.db.voteOrmer.CountByCourseworkID(courseworkID); err != nil {
		return 0, err
	}
	return count, nil
}

func (m *module) VoteCountByUserIDJoinCourseworkType(userID, courseworkType string) (int64, error) {
	var err error
	var count int64
	if count, err = m.db.voteOrmer.CountByUserIDJoinCourseworkType(userID, courseworkType); err != nil {
		return 0, err
	}
	return count, nil
}

func (m *module) VoteHasVoted(userID, courseworkID string) (bool, error) {
	var err error
	if _, err = m.db.voteOrmer.GetOneByUserIDAndCourseworkID(userID, courseworkID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (m *module) VoteUnvote(userID, courseworkID string) error {
	return m.db.voteOrmer.DeleteByUserIDAndCourseworkID(userID, courseworkID)
}

func (m *module) VoteGetVotedProjects(userID string) ([]dtos.Project, error) {
	var err error
	var projectsRaw []models.Project
	if projectsRaw, err = m.db.projectOrmer.GetManyByUserIDJoinVote(userID); err != nil {
		return nil, err
	}
	projects := make([]dtos.Project, len(projectsRaw))
	for i, project := range projectsRaw {
		projects[i] = dtos.Project{
			ID:          project.CourseworkID,
			CourseID:    project.Coursework.CourseID,
			Name:        project.Name,
			Team:        project.Team,
			Description: project.Description,
			Thumbnail:   project.Thumbnail,
			Link:        project.Link,
			Video:       project.Video,
			Field:       project.Field,
			Active:      project.Active,
			Metadata:    project.Metadata,
			CreatedAt:   project.CreatedAt,
		}
	}
	return projects, nil
}

func (m *module) VoteGetVotedBlogs(userID string) ([]dtos.Blog, error) {
	var err error
	var blogsRaw []models.Blog
	if blogsRaw, err = m.db.blogOrmer.GetManyByUserIDJoinVote(userID); err != nil {
		return nil, err
	}
	blogs := make([]dtos.Blog, len(blogsRaw))
	for i, blog := range blogsRaw {
		blogs[i] = dtos.Blog{
			ID:        blog.CourseworkID,
			CourseID:  blog.Coursework.CourseID,
			Title:     blog.Title,
			Author:    blog.Author,
			Link:      blog.Link,
			Category:  blog.Category,
			CreatedAt: blog.CreatedAt,
		}
	}
	return blogs, nil
}
