package requests

import (
	"log"
	"strconv"

	"github.com/PaulicStudios/42-cheat-alert/apimodels"
)

func GetTeamDetails(teamID int) *apimodels.Team {
	var teamDetails apimodels.Team

	url := "https://api.intra.42.fr/v2/teams/" + strconv.Itoa(teamID)
	_, err := request(url, &teamDetails)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &teamDetails
}
