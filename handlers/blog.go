package handlers

import (
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
