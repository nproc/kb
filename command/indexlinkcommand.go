package command

import (
	"strings"

	"github.com/nproc/kb/url"
	"github.com/nproc/kb/user"
)

type IndexLinkCommand interface {
	URL() url.URL
	Tags() []string
	RequestedBy() user.User
}

type indexLinkCommand struct {
	url  url.URL
	user user.User
	tags []string
}

func (c indexLinkCommand) URL() url.URL {
	return c.url
}

func (c indexLinkCommand) Tags() []string {
	return c.tags
}

func (c indexLinkCommand) RequestedBy() user.User {
	return c.user
}

func NewIndexLinkCommandFromUserCommandString(
	user user.User,
	commandString string,
) (IndexLinkCommand, error) {
	parts := strings.SplitN(commandString, " ", 2)

	if len(parts) != 2 {
	}

	url, err := url.New(parts[0])

	if err != nil {
		return nil, err
	}

	// NOTE: need tags pre-validation
	tags := strings.Split(parts[1], " ")

	return indexLinkCommand{
		url:  url,
		user: user,
		tags: tags,
	}, nil
}
