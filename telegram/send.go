package telegram

import "gopkg.in/telebot.v3"

func SendMsgToMe(msg string) {
	b.Send(&telebot.Chat{ID: 379420949}, msg)
}
