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
	if db.DeleteTUser(args[0]) {
		return c.Send("User " + args[0] + " deleted")
	} else if db.AddTUser(args[0]) {
		return c.Send("User " + args[0] + " added")
	} else {
		return c.Send("Problem with adding/deleting user")
	}
}

func toggleAdmin(c tele.Context) error {
	args := c.Args()
	if len(args) != 1 {
		return c.Send("Usage: /toggleadmin <telegram_username>")
	}
	role := db.ToggleAdmin(args[0])
	if role == models.ROLE_ADMIN {
		return c.Send("User " + args[0] + " is now an admin")
	} else if role == models.ROLE_USER {
		return c.Send("User " + args[0] + " is now user")
	} else {
		return c.Send("User " + args[0] + " not found")
	}
}
