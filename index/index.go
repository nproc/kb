package index

import "github.com/nproc/kb/link"

type Index interface {
	Add(link.Link) error
}
