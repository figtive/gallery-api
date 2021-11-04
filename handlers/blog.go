package handlers

import (
	"math/rand"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"
)

func (m *module) BlogInsert(blogInsert dtos.BlogInsert, classID string) (id string, err error) {
	var courseworkID string
	if courseworkID, err = Handler.CourseworkInsert(classID); err != nil {
		return
	}

	if id, err = m.db.blogOrmer.Insert(models.Blog{
		CourseworkID: courseworkID,
		Title:        blogInsert.Title,
		Link:         blogInsert.Link,
		Category:     blogInsert.Category,
		Author:       blogInsert.Author,
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
			ID:       blog.CourseworkID,
			Author:   blog.Author,
			Title:    blog.Title,
			Link:     blog.Link,
			Category: blog.Category,
			CourseId: blog.CourseworkID,
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
		ID:       blogRaw.CourseworkID,
		Author:   blogRaw.Author,
		Title:    blogRaw.Title,
		Link:     blogRaw.Link,
		Category: blogRaw.Category,
		CourseId: blogRaw.CourseworkID,
	}
	return
}
