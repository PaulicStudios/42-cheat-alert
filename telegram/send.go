package telegram

import (
	"github.com/PaulicStudios/42-cheat-alert/apimodels"
	"gopkg.in/telebot.v3"
)

func SendMsgToMe(msg string) {
	b.Send(&telebot.Chat{ID: 379420949}, msg)
}

func SendUpdateMsgs(user *apimodels.User, team *apimodels.Teams) {
	// msg := "Updated team history for team " + team.Name + " with final mark " + team.FinalMark + " for user " + user.Login + " in project " + team.ProjectID
	// SendMsgToMe(msg)
}
