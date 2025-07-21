package main

import (
	"github.com/joho/godotenv"
	"github.com/rizarakhmat/ConcertPing/bot"
	"github.com/rizarakhmat/ConcertPing/scheduler"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	bot.InitBot(botToken)
	bot.RegisterHandlers()
	scheduler.Start()

	bot.Bot.Start()
}
