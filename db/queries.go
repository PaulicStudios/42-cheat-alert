package db

import "github.com/PaulicStudios/42-cheat-alert/models"

func GetAllCheaterTeams() []models.Team {
	var teams []models.Team
	db.Where("final_mark = ?", -42).Find(&teams).Joins("Users")
	return teams
}
