package url_test

import (
	"testing"

	"github.com/nproc/kb/url"
	"github.com/stretchr/testify/assert"
)

func TestNewOrNil(t *testing.T) {
	assert.Nil(t, url.NewOrNil("http"))
	assert.NotNil(t, url.NewOrNil("http://google.com"))
	assert.NotNil(t, url.NewOrNil("https://google.com"))
}

func TestNewURL(t *testing.T) {
	u, err := url.New("https://google.com")

	assert.Nil(t, err)
	assert.Equal(t, "https://google.com", u.URLString())
	assert.Equal(t, "https://google.com", u.String())
}

func TestNewInvalidURL(t *testing.T) {
	u, err := url.New("https:\\google.com")

	assert.Nil(t, u)
	assert.Equal(t, url.ErrInvalidURL, err)
}
