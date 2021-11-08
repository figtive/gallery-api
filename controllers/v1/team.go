package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
)

func POSTTeam(c *gin.Context) {
	var err error

	var teamInsert dtos.TeamInsert
	if err = c.ShouldBindJSON(&teamInsert); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{Error: err.Error()})
		return
	}

	var teamId string
	// if teamId, err = handlers.Handler.TeamInsert(teamInsert); err != nil {
	// 	c.JSON(http.StatusInternalServerError, dtos.Response{Error: err})
	// 	return
	// }

	c.JSON(http.StatusOK, dtos.Response{
		Code: http.StatusOK,
		Data: dtos.Team{
			ID:   teamId,
			Name: teamInsert.Name,
		},
	})
}
