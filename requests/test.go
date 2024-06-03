package requests

import (
	"log"

	"github.com/PaulicStudios/42-cheat-alert/apimodels"
)

func TestGetTeamDetails(teamID int) *apimodels.Team {
	var teamDetails apimodels.Team

	url := "https://api.intra.42.fr/v2/teams/4813323"
	_, err := request(url, &teamDetails)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &teamDetails
}
