package main

import (
	"log"
	"os"
	"time"

	"github.com/nproc/kb/command"
	"github.com/nproc/kb/user"
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

	commandBus := command.NewCommandBus()

	for message := range messages {
		user := user.FromTelegramMessage(message)

		commandBus(
			user,
			message.Text,
		)
	}

	return nil
}
