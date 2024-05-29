package cheatalert

import (
	"log"
	"strconv"

	"github.com/PaulicStudios/42-cheat-alert/apimodels"
	"github.com/PaulicStudios/42-cheat-alert/db"
	"github.com/PaulicStudios/42-cheat-alert/requests"
	"github.com/PaulicStudios/42-cheat-alert/telegram"
)

func UpdateProjects() {
	nextPage := true
	page := 0
	for nextPage {
		page++
		var projectUsers *apimodels.ProjectsUsers
		projectUsers, nextPage = requests.GetRecentProjectUsers(page)
		for _, projectUser := range *projectUsers {
			db.SaveApiUser(&projectUser.User)
			for _, team := range projectUser.Teams {
				db.SaveApiTeam(&projectUser.User, &team)
				if db.UpdateTeamHistory(&team) {
					telegram.SendMsgToMe("Updated team history for team " + team.Name + " with final mark " + strconv.Itoa(team.FinalMark) + " for user " + projectUser.User.Login + " in project " + team.ProjectGitlabPath)
					log.Println("Updated team history for team", team.Name, "with final mark", team.FinalMark, "for user", projectUser.User.Login, "in project", team.ProjectID)
				}
			}
		}
	}
	log.Println("Updated projects with ", page, " requests")
}
