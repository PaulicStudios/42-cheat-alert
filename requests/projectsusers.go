package requests

import (
	"log"
	"strconv"
	"time"

	"github.com/PaulicStudios/42-cheat-alert/apimodels"
)

func GetRecentProjectUsers(page int) (*apimodels.ProjectsUsers, bool) {
	projects_users := apimodels.ProjectsUsers{}
	now := time.Now().Format("2006-01-02T15:04:05.000Z")
	yearAgo := time.Now().AddDate(0, 0, -1).Format("2006-01-02T15:04:05.000Z")
	url := "https://api.intra.42.fr/v2/projects_users" +
		// "?filter[final_mark]=-42" +
		"?filter[campus]=39" +
		"&range[updated_at]=" + yearAgo + "," + now +
		"&page[number]=" + strconv.Itoa(page) + "&page[size]=100"
	responce, err := request(url, &projects_users)
	if err != nil {
		log.Println(err)
		return nil, false
	}
	return &projects_users, hasNextPage(responce)
}
