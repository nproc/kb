package link_test

import (
	"testing"

	"github.com/nproc/kb/link"
	"github.com/nproc/kb/url"
	"github.com/stretchr/testify/assert"
)

func TestNewLink(t *testing.T) {
	url, err := url.New("https://google.com")

	assert.Nil(t, err)

	link := link.New(url)

	assert.NotEmpty(t, link.ID().String())
	assert.Equal(t, "https://google.com", link.URL().String())
}
