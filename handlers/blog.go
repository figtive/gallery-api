package handlers

import (
	"math/rand"
	"time"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/utils"
)

func (m *module) BlogInsert(blogInsert dtos.BlogInsert, classID string) (id string, err error) {
	var courseworkID string
	if courseworkID, err = Handler.CourseworkInsert(classID); err != nil {
		return
	}

	if id, err = m.db.blogOrmer.Insert(models.Blog{
		CourseworkID: courseworkID,
		Title:        blogInsert.Title,
		Author:       blogInsert.Author,
		Link:         blogInsert.Link,
		Category:     blogInsert.Category,
	}); err != nil {
		return
	}
	return
}

func (m *module) BlogGetMany(skip int, limit int) (blogs []dtos.Blog, err error) {
	var blogsRaw []models.Blog
	if blogsRaw, err = m.db.blogOrmer.GetMany(skip, limit); err != nil {
		return
	}
	blogs = make([]dtos.Blog, len(blogsRaw))
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
	return
}

func (m *module) BlogGetOne(id string) (blog dtos.Blog, err error) {
	var blogRaw models.Blog
	if blogRaw, err = m.db.blogOrmer.GetOneByCourseworkID(id); err != nil {
		return
	}
	blog = dtos.Blog{
		ID:        blogRaw.CourseworkID,
		CourseID:  blogRaw.Coursework.CourseID,
		Title:     blogRaw.Title,
		Author:    blogRaw.Author,
		Link:      blogRaw.Link,
		Category:  blogRaw.Category,
		CreatedAt: blogRaw.CreatedAt,
	}
	return
}

func (m *module) BlogGetManyByCourseIDInCurrentTerm(courseID string) ([]dtos.Blog, error) {
	var err error
	var blogsRaw []models.Blog
	if blogsRaw, err = m.db.blogOrmer.GetManyByCourseIDAndTerm(courseID, utils.TimeToTermTime(time.Now()), utils.NextTermTime(time.Now())); err != nil {
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
