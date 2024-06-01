package telegram

import (
	"github.com/PaulicStudios/42-cheat-alert/db"
	tele "gopkg.in/telebot.v3"
)

func notifyCommand(c tele.Context) error {
	if db.ToggleNotify(c.Sender().ID) {
		return c.Send("Notifications enabled")
	} else {
		return c.Send("Notifications disabled")
	}
}
