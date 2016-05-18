package link

import (
	"github.com/nproc/kb/url"
	"github.com/satori/go.uuid"
)

type Link interface {
	ID() uuid.UUID
	URL() url.URL
}

type link struct {
	id  uuid.UUID
	url url.URL
}

func (l link) ID() uuid.UUID {
	return l.id
}

func (l link) URL() url.URL {
	return l.url
}

func New(url url.URL) Link {
	return &link{
		id:  uuid.NewV4(),
		url: url,
	}
}
