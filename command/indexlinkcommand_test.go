package command_test

import (
	"testing"

	"github.com/nproc/kb/command"
	"github.com/nproc/kb/user"
	"github.com/stretchr/testify/assert"
)

func TestNewIndexLinkCommandFromUserCommandString(t *testing.T) {
	command, err := command.NewIndexLinkCommandFromUserCommandString(
		user.New("Andrey", "CentaurWarchief"),
		"https://github.com #github #gh #os",
	)

	assert.Nil(t, err)
	assert.NotNil(t, command)
	assert.Equal(t, "https://github.com", command.URL().String())
	assert.Len(t, command.Tags(), 3)
}
