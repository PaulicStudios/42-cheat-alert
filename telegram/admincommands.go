package telegram

import (
	"github.com/PaulicStudios/42-cheat-alert/db"
	"github.com/PaulicStudios/42-cheat-alert/models"
	tele "gopkg.in/telebot.v3"
)

func toggleUser(c tele.Context) error {
	args := c.Args()
	if len(args) != 1 {
		return c.Send("Usage: /toggleuser <telegram_username>")
	}
	chat, err := c.Bot().ChatByUsername(args[0])
	if err != nil {
		return c.Send("Invalid telegram_username")
	}
	if db.DeleteTUser(chat.ID) {
		return c.Send("User " + chat.FirstName + " deleted")
	} else if db.AddTUser(chat.ID) {
		return c.Send("User " + chat.FirstName + " added")
	} else {
		return c.Send("Problem with adding/deleting user")
	}
}

func toggleAdmin(c tele.Context) error {
	args := c.Args()
	if len(args) != 1 {
		return c.Send("Usage: /toggleadmin <telegram_username>")
	}
	chat, err := c.Bot().ChatByUsername(args[0])
	if err != nil {
		return c.Send("Invalid telegram_username")
	}
	role := db.ToggleAdmin(chat.ID)
	if role == models.ROLE_ADMIN {
		return c.Send("User " + chat.FirstName + " is now an admin")
	} else if role == models.ROLE_USER {
		return c.Send("User " + chat.FirstName + " is now user")
	} else {
		return c.Send("User " + chat.FirstName + " not found")
	}
}
