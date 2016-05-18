package index

import (
	"github.com/nproc/kb/link"
	"github.com/satori/go.uuid"
)

type Index interface {
	Add(link.Link) error
	Remove(link.Link) error
	Get(uuid.UUID) (link.Link, error)
}
