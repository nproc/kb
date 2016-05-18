package link

import (
	"time"

	"github.com/nproc/kb/url"
	"github.com/nproc/kb/user"
	"github.com/satori/go.uuid"
)

type Link interface {
	ID() uuid.UUID
	URL() url.URL
	SharedBy() user.User
	SharedAt() time.Time
}

type link struct {
	id       uuid.UUID
	url      url.URL
	user     user.User
	sharedAt time.Time
}

func (l link) ID() uuid.UUID {
	return l.id
}

func (l link) URL() url.URL {
	return l.url
}

func (l link) SharedBy() user.User {
	return l.user
}

func (l link) SharedAt() time.Time {
	return l.sharedAt
}

func New(url url.URL, user user.User) Link {
	return &link{
		id:       uuid.NewV4(),
		url:      url,
		user:     user,
		sharedAt: time.Now(),
	}
}
