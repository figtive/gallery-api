package handlers

import (
	"math/rand"

	"gorm.io/gorm"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"
)

func (m *module) BookmarkInsert(bookmark dtos.Bookmark) (string, error) {
	var err error
	var id string
	if id, err = m.db.bookmarkOrmer.Insert(models.Bookmark{
		UserID:       bookmark.UserID,
		CourseworkID: bookmark.CourseworkID,
	}); err != nil {
		return "", err
	}
	return id, nil
}

func (m *module) BookmarkHasMarked(bookmark dtos.Bookmark) (bool, error) {
	var err error
	if _, err = m.db.bookmarkOrmer.GetOneByUserIDAndCourseworkID(bookmark.UserID, bookmark.CourseworkID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (m *module) BookmarkDelete(bookmark dtos.Bookmark) error {
	var err error
	if err = m.db.bookmarkOrmer.DeleteByUserIDAndCourseworkID(bookmark.UserID, bookmark.CourseworkID); err != nil {
		return err
	}
	return nil
}

func (m *module) BookmarkGetManyBlogByUserID(userID string) ([]dtos.Blog, error) {
	var err error
	var blogsRaw []models.Blog
	if blogsRaw, err = m.db.blogOrmer.GetManyBookmarkByUserID(userID); err != nil {
		return nil, err
	}
	blogs := make([]dtos.Blog, len(blogsRaw))
	for i, j := range rand.Perm(len(blogsRaw)) {
		blog := blogsRaw[j]
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

func (m *module) BookmarkGetManyProjectByUserID(userID string) ([]dtos.Project, error) {
	var err error
	var projectsRaw []models.Project
	if projectsRaw, err = m.db.projectOrmer.GetManyBookmarkByUserID(userID); err != nil {
		return nil, err
	}
	projects := make([]dtos.Project, len(projectsRaw))
	for i, j := range rand.Perm(len(projectsRaw)) {
		project := projectsRaw[j]
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
