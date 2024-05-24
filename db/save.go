package db

import "github.com/PaulicStudios/42-cheat-alert/models"

func SaveUser(user models.User) {
	db.Save(&user)
}
