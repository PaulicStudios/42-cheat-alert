package db

import (
	"strings"

	"github.com/PaulicStudios/42-cheat-alert/models"
	tele "gopkg.in/telebot.v3"
)

func UserExists(tUser *tele.User) bool {
	var count int64
	db.Model(&models.TUser{}).Where("t_user_id = ?", tUser.ID).Count(&count)
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

func StartTUser(tUser *tele.User) bool {
	if UserExists(tUser) {
		return true
	}
	res := db.Where("user_name = ?", strings.ToLower(tUser.Username)).First(&models.TUser{}).Update("t_user_id", tUser.ID)
	return res.RowsAffected != 0
}

func ToggleNotify(tuserID int64) (bool, error) {
	var user models.TUser
	res := db.Where("t_user_id = ?", tuserID).First(&user).Select("id", "notify")
	if res.Error != nil {
		return false, res.Error
	}
	user.Notify = !user.Notify
	res = db.Save(&user)
	if res.Error != nil {
		return false, res.Error
	}
	return user.Notify, nil
}

func ToggleAdmin(userName string) int {
	var user models.TUser
	userName = strings.ToLower(userName)
	res := db.Where("user_name = ?", userName).First(&user).Select("role")
	if res.RowsAffected == 0 {
		return -1
	}
	if user.Role == models.ROLE_ADMIN {
		user.Role = models.ROLE_USER
	} else {
		user.Role = models.ROLE_ADMIN
	}
	db.Save(&user)
	return user.Role
}

func AddTUser(userName string) bool {
	userName = strings.ToLower(userName)
	res := db.Where("user_name = ?", userName).FirstOrCreate(&models.TUser{UserName: userName}).Select("id")
	return res.RowsAffected != 0
}

func DeleteTUser(userName string) bool {
	userName = strings.ToLower(userName)
	res := db.Where("user_name = ?", userName).Delete(&models.TUser{})
	return res.RowsAffected != 0
}
