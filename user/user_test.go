package user_test

import (
	"testing"

	"github.com/nproc/kb/user"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user := user.New("Tarcísio Gruppi", "txgruppi")

	assert.NotNil(t, user)
	assert.Equal(t, "Tarcísio Gruppi", user.FullName())
	assert.Equal(t, "txgruppi", user.Username())
}
