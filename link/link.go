package link

import (
	"encoding/json"
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
	Tags() []string
	json.Marshaler
}

type link struct {
	id       uuid.UUID
	url      url.URL
	sharedBy user.User
	sharedAt time.Time
	tags     []string
}

func (l link) ID() uuid.UUID {
	return l.id
}

func (l link) URL() url.URL {
	return l.url
}

func (l link) SharedBy() user.User {
	return l.sharedBy
}

func (l link) SharedAt() time.Time {
	return l.sharedAt
}

func (l link) Tags() []string {
	return l.tags
}

func (l link) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID       string    `json:"id"`
		URL      string    `json:"url"`
		SharedBy user.User `json:"shared_by"`
		SharedAt time.Time `json:"shared_at"`
		Tags     []string  `json:"tags"`
	}{
		l.ID().String(),
		l.url.String(),
		l.sharedBy,
		l.sharedAt,
		l.tags,
	})
}

func New(url url.URL, user user.User, tags []string) Link {
	// NOTE: filter duplicate tags
	// NOTE: make sure that have at least one tag
	// NOTE: should have at most `N` (to be decided) tags

	return &link{
		id:       uuid.NewV4(),
		url:      url,
		sharedBy: user,
		sharedAt: time.Now(),
		tags:     tags,
	}
}
