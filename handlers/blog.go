package handlers

import (
	"math/rand"
	"time"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/constants"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/utils"
)

func (m *module) BlogInsert(blogInsert dtos.BlogInsert, classID string) (string, error) {
	var err error

	var courseworkID string
	if courseworkID, err = Handler.CourseworkInsert(classID, constants.CourseworkTypeBlog); err != nil {
		return "", err
	}
	var id string
	if id, err = m.db.blogOrmer.Insert(models.Blog{
		CourseworkID: courseworkID,
		Title:        blogInsert.Title,
		Author:       blogInsert.Author,
		Link:         blogInsert.Link,
		Category:     blogInsert.Category,
	}); err != nil {
		return "", err
	}
	return id, nil
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

func (m *module) BlogGetManyByCourseIDInCurrentTerm(courseID string, currentOnly bool) ([]dtos.Blog, error) {
	var err error
	var startTime, endTime time.Time
	if currentOnly {
		startTime = utils.TimeToTermTime(time.Now())
		endTime = utils.NextTermTime(time.Now())
	} else {
		startTime = time.Unix(-2208988800, 0)
		endTime = startTime.Add(1<<63 - 1)
	}
	var blogsRaw []models.Blog
	if blogsRaw, err = m.db.blogOrmer.GetManyByCourseIDAndTerm(courseID, startTime, endTime); err != nil {
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
