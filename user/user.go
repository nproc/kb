package user

import (
	"fmt"
	"strings"

	"github.com/tucnak/telebot"
)

type User interface {
	FullName() string
	Username() string
}

type user struct {
	fullName string
	username string
}

func (u user) FullName() string {
	return u.fullName
}

func (u user) Username() string {
	return u.username
}

func FromTelegramMessage(message telebot.Message) User {
	fullName := strings.TrimSpace(
		fmt.Sprintf(
			"%s %s",
			message.Sender.FirstName,
			message.Sender.LastName,
		),
	)

	return &user{
		fullName: fullName,
		username: message.Sender.Username,
	}
}

func New(fullName, username string) User {
	return &user{
		fullName,
		username,
	}
}
