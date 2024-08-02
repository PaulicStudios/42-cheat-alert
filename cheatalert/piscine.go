package cheatalert

import (
	"log"
	"strconv"

	"github.com/PaulicStudios/42-cheat-alert/apimodels"
	"github.com/PaulicStudios/42-cheat-alert/db"
	"github.com/PaulicStudios/42-cheat-alert/requests"
)

func GetPisciners() {
	nextPage := true
	page := 0
	for nextPage {
		page++
		var users *apimodels.UserList
		users, nextPage = requests.GetUsersFromPiscine(page)
		for _, user := range *users {
			userDetails := requests.GetUser(strconv.Itoa(user.ID))
			if userDetails == nil {
				log.Println("Failed to get user details for", user.Login)
				continue
			}
			db.SaveApiDetailUser(userDetails)
		}
	}
}
