package requests

import (
	"log"
	"strconv"

	"github.com/PaulicStudios/42-cheat-alert/apimodels"
)

func GetTeamDetails(teamID int) *apimodels.Team {
	teamIDstr := strconv.Itoa(teamID)
	var teamDetails apimodels.Team

	url := "https://api.intra.42.fr/v2/teams/" + teamIDstr
	_, err := request(url, &teamDetails)
	if err != nil {
		log.Println(err, "for team", teamIDstr)
		return nil
	}
	return &teamDetails
}
