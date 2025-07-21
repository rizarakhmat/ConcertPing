package scheduler

import (
	"fmt"

	"github.com/rizarakhmat/ConcertPing/bot"
	"github.com/rizarakhmat/ConcertPing/service"
	"github.com/rizarakhmat/ConcertPing/storage"
	"github.com/robfig/cron/v3"
	"gopkg.in/telebot.v3"
)

func Start() {
	c := cron.New()
	c.AddFunc("0 9 * * 1", sendWeeklyUpdates) // every Monday 9AM
	c.Start()
}

func sendWeeklyUpdates() {
	users := storage.GetAllUsers()
	for _, user := range users {
		events, err := service.GetConcerts(user.Artist, user.Latitude, user.Longitude)
		if err != nil || len(events) == 0 {
			bot.Bot.Send(&telebot.User{ID: int64(user.UserID)}, "No concerts found this week.")
			continue
		}

		msg := "🎶 Upcoming concerts for " + user.Artist + ":\n"
		for _, e := range events {
			msg += fmt.Sprintf("• %s\n%s\n\n", e.Name, e.URL)
		}
		bot.Bot.Send(&telebot.User{ID: int64(user.UserID)}, msg)
	}
}
