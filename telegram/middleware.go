package telegram

import (
	"github.com/PaulicStudios/42-cheat-alert/db"
	tele "gopkg.in/telebot.v3"
)

func isAllowed() tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if !db.UserExists(c.Sender().ID) {
				return c.Send("You are not allowed to use this bot")
			}
			return next(c)
		}
	}
}

func isAdmin() tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if !db.UserIsAdmin(c.Sender().ID) {
				return c.Send("You are not allowed to use this command")
			}
			return next(c)
		}
	}
}
