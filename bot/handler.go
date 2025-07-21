package bot

import (
	"github.com/rizarakhmat/ConcertPing/storage"
	"gopkg.in/telebot.v3"
)

func RegisterHandlers() {
	Bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Hi! Please share your location.")
	})

	Bot.Handle(telebot.OnLocation, func(c telebot.Context) error {
		loc := c.Message().Location
		storage.SaveUserLocation(c.Sender().ID, float32(loc.Lat), float32(loc.Lng))
		return c.Send("Great! Now tell me your favorite artist.")
	})

	Bot.Handle(telebot.OnText, func(c telebot.Context) error {
		storage.SaveUserArtist(c.Sender().ID, c.Text())
		return c.Send("Thanks! You'll get weekly concert updates every Monday.")
	})
}
