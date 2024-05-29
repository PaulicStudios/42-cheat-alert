package cheatalert

import (
	"strconv"

	"github.com/PaulicStudios/42-cheat-alert/db"
	"github.com/PaulicStudios/42-cheat-alert/requests"
	"github.com/PaulicStudios/42-cheat-alert/telegram"
)

func UpdateProjects() {
	projectUsers := requests.GetRecentProjectUsers()
	for _, projectUser := range projectUsers {
		db.SaveApiUser(&projectUser.User)
		for _, team := range projectUser.Teams {
			db.SaveApiTeam(&projectUser.User, &team)
			if db.UpdateTeamHistory(&team) {
				telegram.SendMsgToMe("Added new team history for team " + team.Name + " with final mark " + strconv.Itoa(team.FinalMark) + " for user " + projectUser.User.Login)
			}
		}
	}
}
