package cheatalert

import (
	"log"

	"github.com/PaulicStudios/42-cheat-alert/db"
	"github.com/PaulicStudios/42-cheat-alert/requests"
)

func UpdateProjects() {
	projectUsers := requests.GetRecentProjectUsers()
	for _, projectUser := range projectUsers {
		db.SaveApiUser(&projectUser.User)
		for _, team := range projectUser.Teams {
			db.SaveApiTeam(&projectUser.User, &team)
			if db.UpdateTeamHistory(&team) {
				log.Println("User: ", projectUser.User.Login)
				log.Println("Project: ", team.ProjectID)
				log.Println("Team: ", team.ID)
				log.Println("Mark: ", team.FinalMark)
			}
		}
	}
}
