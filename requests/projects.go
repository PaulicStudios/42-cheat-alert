package requests

import (
	"log"

	"github.com/PaulicStudios/42-cheat-alert/models"
)

func GetRecentProjectUsers() models.ProjectsUsers {
	resp := models.ProjectsUsers{}
	err := request("https://api.intra.42.fr/v2/projects_users?filter[project_id]=1&filter[status]=finished", &resp)
	if err != nil {
		log.Println(err)
		return nil
	}
	return resp
}
