package command_test

import (
	"testing"

	"github.com/nproc/kb/command"
	"github.com/stretchr/testify/assert"
)

func TestNewIndexLinkCommandFromCommandString(t *testing.T) {
	command, err := command.NewIndexLinkCommandFromCommandString(
		"https://github.com #github #gh #os",
	)

	assert.Nil(t, err)
	assert.NotNil(t, command)
	assert.Equal(t, "https://github.com", command.URL().String())
	assert.Len(t, command.Tags(), 3)
}
