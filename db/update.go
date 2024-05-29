package db

import (
	"github.com/PaulicStudios/42-cheat-alert/apimodels"
	"github.com/PaulicStudios/42-cheat-alert/models"
)

func UpdateTeamHistory(team *apimodels.Teams) bool {
	var TeamHistory models.TeamHistory = models.TeamHistory{
		TeamID: team.ID,
	}
	result := db.Order("created_at desc").Where("team_id = ?", team.ID).First(&TeamHistory)
	if result.Error != nil {
		TeamHistory = models.TeamHistory{
			TeamID:    team.ID,
			FinalMark: team.FinalMark,
		}
		db.Create(&TeamHistory)
		return true
	}
	if TeamHistory.FinalMark != team.FinalMark {
		TeamHistory = models.TeamHistory{
			TeamID:    team.ID,
			FinalMark: team.FinalMark,
		}
		db.Create(&TeamHistory)
	}
	return TeamHistory.FinalMark != team.FinalMark
}
