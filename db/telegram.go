package db

import "github.com/PaulicStudios/42-cheat-alert/models"

func ToggleNotify(tuserID int64) bool {
	var user models.TUser
	db.Where("t_user_id = ?", tuserID).First(&user).Select("notify")
	user.Notify = !user.Notify
	db.Save(&user)
	return user.Notify
}

func UserExists(tuserID int64) bool {
	var count int64
	db.Model(&models.TUser{}).Where("t_user_id = ?", tuserID).Count(&count)
	return count > 0
}

func UserIsAdmin(tuserID int64) bool {
	var user models.TUser
	db.Where("t_user_id = ?", tuserID).First(&user).Select("role")
	return user.Role == models.ROLE_ADMIN
}

func AllNotifyUsers() []models.TUser {
	var users []models.TUser
	db.Where("notify = ?", true).Find(&users).Select("t_user_id")
	return users
}
