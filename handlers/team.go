package handlers

import (
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"
)

func (m *module) TeamInsert(teamInfo dtos.TeamInsert) (id string, err error) {
	if id, err = m.db.teamOrmer.Insert(models.Team{
		Name:      teamInfo.Name,
		ClassID:   teamInfo.ClassID,
		ProjectID: teamInfo.ProjectID,
	}); err != nil {
		return
	}
	return
}
