package command

import (
	"strings"

	"github.com/nproc/kb/url"
)

type IndexLinkCommand interface {
	URL() url.URL
	Tags() []string
}

type indexLinkCommand struct {
	url  url.URL
	tags []string
}

func (c indexLinkCommand) URL() url.URL {
	return c.url
}

func (c indexLinkCommand) Tags() []string {
	return c.tags
}

func NewIndexLinkCommandFromCommandString(
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
		tags: tags,
	}, nil
}
