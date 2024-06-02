package telegram

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

var b *tele.Bot

func Init_telegram() {
	pref := tele.Settings{
		Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	var err error
	b, err = tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Use(isAllowed())
	b.Handle("/notify", notifyCommand)

	adminOnly := b.Group()
	adminOnly.Use(isAdmin())
	adminOnly.Handle("/toggleuser", toggleUser)
	adminOnly.Handle("/toggleadmin", toggleAdmin)

	b.Start()
}
