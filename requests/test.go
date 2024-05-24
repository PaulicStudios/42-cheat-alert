package requests

import (
	"log"

	"github.com/PaulicStudios/42-cheat-alert/models"
)

func Test() {
	list := models.ProjectsUsers{}
	err := request("https://api.intra.42.fr/v2/projects_users?filter[project_id]=1&filter[status]=finished", &list)
	if err != nil {
		log.Println(err)
		return
	}
	for _, projectUser := range list {
		println(projectUser.User.Login, projectUser.Project.Name, projectUser.Status)
	}
}
