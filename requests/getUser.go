package requests

import (
	"log"
	"strconv"

	"github.com/PaulicStudios/42-cheat-alert/apimodels"
)

func GetUser(identifier string) *apimodels.UserDetailed {
	user := apimodels.UserDetailed{}
	url := "https://api.intra.42.fr/v2/users/" + identifier
	_, err := request(url, &user)
	if err != nil {
		log.Println(err, "for getting User")
		return nil
	}
	return &user
}

func GetUsersFromPiscine(page int) (*apimodels.UserList, bool) {
	userList := apimodels.UserList{}

	url := "https://api.intra.42.fr/v2/campus/39/users" +
		"?filter[pool_month]=august" +
		"&filter[pool_year]=2024" +
		"&page[number]=" + strconv.Itoa(page) + "&page[size]=100"

	request, err := request(url, &userList)
	if err != nil {
		log.Println(err, "for getting Users")
		return nil, false
	}
	return &userList, hasNextPage(request)
}
