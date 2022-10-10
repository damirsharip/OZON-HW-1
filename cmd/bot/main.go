package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"HW-1/config"
	"HW-1/internal/commander"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		log.Panic(err)
	}

	cmd, err := commander.Init(bot)
	if err != nil {
		log.Panic(err)
	}
	cmd.Run()

}
