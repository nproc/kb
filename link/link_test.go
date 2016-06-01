package link_test

import (
	"testing"

	"github.com/nproc/kb/link"
	"github.com/nproc/kb/url"
	"github.com/nproc/kb/user"
	"github.com/stretchr/testify/assert"
)

func TestNewLink(t *testing.T) {
	url, err := url.New("https://google.com")

	assert.Nil(t, err)

	user := user.New("Tarcísio Gruppi", "txgruppi")

	link := link.New(url, user, []string{"#google"})

	assert.NotEmpty(t, link.ID().String())
	assert.Equal(t, "https://google.com", link.URL().String())
	assert.Equal(t, user, link.SharedBy())
	assert.EqualValues(t, []string{"#google"}, link.Tags())
}
