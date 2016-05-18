package main

import (
	"log"
	"os"
	"time"

	"github.com/tucnak/telebot"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	bot, err := telebot.NewBot(os.Getenv("TELEGRAM_KB_BOT_TOKEN"))

	if err != nil {
		return err
	}

	messages := make(chan telebot.Message)

	bot.Listen(messages, 1*time.Second)

	return nil
}
