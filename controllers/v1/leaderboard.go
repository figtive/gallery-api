package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/handlers"
)

func GETProjectLeaderboard(c *gin.Context) {
	var err error

	courseID := c.Param("id")

	var projects []dtos.Project
	if projects, err = handlers.Handler.LeaderboardProject(time.Now(), courseID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}
