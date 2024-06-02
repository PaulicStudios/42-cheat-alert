package telegram

import (
	"log"

	"github.com/PaulicStudios/42-cheat-alert/db"
	tele "gopkg.in/telebot.v3"
)

func notifyCommand(c tele.Context) error {
	notify, err := db.ToggleNotify(c.Sender().ID)
	if err != nil {
		log.Println(err)
		return c.Send("Error updating setting")
	}
	if notify {
		return c.Send("Notifications enabled")
	} else {
		return c.Send("Notifications disabled")
	}
}

func startCommand(c tele.Context) error {
	if !db.StartTUser(c.Sender()) {
		return c.Send("You are not allowed to use this bot")
	}
	return c.Send("Welcome to 42 Cheat Alert Bot!\n" +
		"Use /notify to start receiving notifications")
}
