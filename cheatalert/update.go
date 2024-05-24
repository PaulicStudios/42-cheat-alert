package cheatalert

import (
	"github.com/PaulicStudios/42-cheat-alert/db"
	"github.com/PaulicStudios/42-cheat-alert/requests"
)

func UpdateProjects() {
	projectUsers := requests.GetRecentProjectUsers()
	for _, projectUser := range projectUsers {
		db.SaveApiUser(&projectUser.User)
		// db.SaveTeams(projectUser.Teams)
	}
}
