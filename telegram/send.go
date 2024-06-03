package telegram

import (
	"strconv"

	"github.com/PaulicStudios/42-cheat-alert/apimodels"
	"github.com/PaulicStudios/42-cheat-alert/db"
	"gopkg.in/telebot.v3"
)

func SendMsg(msg string, id int64) {
	b.Send(&telebot.Chat{ID: id}, msg)
}

func SendUpdateMsgs(user *apimodels.User, team *apimodels.Teams, project *apimodels.Project) {
	var msg string
	if team.FinalMark == -42 {
		msg = "Cheater Detected!\n"
	} else {
		msg = "Detected Change!\n"
	}
	msg += "Team:\n" +
		"- Name: " + team.Name + "\n" +
		"- Users: "
	for ind, user := range team.Users {
		msg += "https://profile.intra.42.fr/users/" + user.Login
		if ind < len(team.Users)-1 {
			msg += ", "
		}
	}
	msg += "\n" +
		"- Final Mark: " + strconv.Itoa(team.FinalMark) + "\n" +
		"- Project: " + project.Name + "\n" +
		"- Status: " + team.Status + "\n" +
		"- closed: " + strconv.FormatBool(team.Closed) + "\n" +
		"- Validated: " + strconv.FormatBool(team.Validated) + "\n" +
		"- Locked: " + strconv.FormatBool(team.Locked) + "\n" +
		"- Locked At: " + team.LockedAt.Local().Format("02.01.2006 15:04") + "\n" +
		"- Created At: " + team.CreatedAt.Local().Format("02.01.2006 15:04")
	for _, tUser := range db.AllNotifyUsers() {
		SendMsg(msg, tUser.TUserID)
	}
}
