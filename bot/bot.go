package bot

import (
	"log"
	"time"

	"gopkg.in/telebot.v3"
)

var Bot *telebot.Bot

func InitBot(token string) {
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	var err error
	Bot, err = telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
}
