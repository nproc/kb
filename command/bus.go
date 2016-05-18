package command

import (
	"log"
	"strings"

	"github.com/nproc/kb/user"
)

// NOTE: can be A LOT improved

func IndexLinkCommandHandler(command interface{}) {
	if _, ok := command.(IndexLinkCommand); !ok {
		return
	}
}

type CommandBus func(user.User, string)

func NewCommandBus() CommandBus {
	var (
		handlers = make(map[string]Handler)
		parsers  = make(map[string]Parser)
	)

	handlers["add"] = IndexLinkCommandHandler

	parsers["add"] = func(user user.User, commandString string) (interface{}, error) {
		return NewIndexLinkCommandFromUserCommandString(
			user,
			commandString,
		)
	}

	return func(user user.User, commandString string) {
		parts := strings.SplitN(commandString, " ", 3)

		if len(parts) != 3 {
			return
		}

		if parts[0] != "/kb" {
			return
		}

		if _, ok := parsers[parts[1]]; !ok {
			return
		}

		if _, ok := handlers[parts[1]]; !ok {
			return
		}

		command, err := parsers[parts[1]](user, parts[2])

		if err != nil {
			log.Println(err)
			return
		}

		handlers[parts[1]](command)
	}
}
