package requests

import (
	"log"
	"time"

	"github.com/PaulicStudios/42-cheat-alert/apimodels"
)

func GetRecentProjectUsers() apimodels.ProjectsUsers {
	resp := apimodels.ProjectsUsers{}
	now := time.Now().Format("2006-01-02T15:04:05.000Z")
	yearAgo := time.Now().AddDate(0, -1, 0).Format("2006-01-02T15:04:05.000Z")
	url := "https://api.intra.42.fr/v2/projects_users" +
		"?filter[final_mark]=-42" +
		"&filter[campus]=39" +
		"&range[updated_at]=" + yearAgo + "," + now +
		"&page[number]=1&page[size]=100"
	err := request(url, &resp)
	if err != nil {
		log.Println(err)
		return nil
	}
	return resp
}
