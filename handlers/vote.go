package handlers

import (
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"
)

func (m *module) VoteInsert(voteInfo dtos.VoteInsert) (string, error) {
	var err error
	var id string
	if id, err = m.db.voteOrmer.Insert(models.Vote{
		UserID:       voteInfo.UserID,
		CourseworkID: voteInfo.CourseworkID,
	}); err != nil {
		return "", err
	}
	return id, nil
}
