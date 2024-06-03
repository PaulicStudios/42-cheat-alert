package cheatalert

import (
	"log"

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
				if db.SaveApiTeam(&projectUser.User, &team) {
					teamDetails := requests.GetTeamDetails(team.ID)
					if teamDetails != nil {
						db.SaveApiTeamDetails(teamDetails)
					}
					telegram.SendUpdateMsgs(&projectUser.User, teamDetails, &projectUser.Project)
					log.Println("Updated team history for team", team.Name, "with final mark", team.FinalMark, "for user", projectUser.User.Login, "in project", team.ProjectID)
				}
			}
		}
	}
	log.Println("Updated projects with ", page, " requests")
}

func UpdateProjectsLastMonth() {
	log.Print("Getting Every Project")
	nextPage := true
	page := 0
	for nextPage {
		print(".")
		page++
		var projectUsers *apimodels.ProjectsUsers
		projectUsers, nextPage = requests.GetRecentProjectUsers(page)
		for _, projectUser := range *projectUsers {
			db.SaveApiUser(&projectUser.User)
			for _, team := range projectUser.Teams {
				teamDetails := requests.GetTeamDetails(team.ID)
				if teamDetails != nil {
					db.SaveApiTeamDetails(teamDetails)
				}
				if db.SaveApiTeam(&projectUser.User, &team) {
					// telegram.SendUpdateMsgs(&projectUser.User, &team, &projectUser.Project)
					log.Println("Updated team history for team", team.Name, "with final mark", team.FinalMark, "for user", projectUser.User.Login, "in project", team.ProjectID)
				}
			}
		}
	}
	log.Println("Updated projects with ", page, " requests")
}
