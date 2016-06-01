package index

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/nproc/kb/link"
)

var (
	BoltLinksBucket = []byte("kb:links")
	BoltTagsBucket  = []byte("kb:tags2link")
)

type Bolt struct {
	db *bolt.DB
}

func (b Bolt) Add(link link.Link) error {
	return b.db.Update(func(tx *bolt.Tx) error {
		serialized, err := json.Marshal(link)

		if err != nil {
			return err
		}

		var (
			tags  = tx.Bucket(BoltTagsBucket)
			links = tx.Bucket(BoltLinksBucket)
		)

		linkID := []byte(link.ID().String())

		for _, tag := range link.Tags() {
			if err := tags.Put([]byte(tag), linkID); err != nil {
				return err
			}
		}

		return links.Put(linkID, serialized)
	})
}

func NewBoltIndex(db *bolt.DB) (Index, error) {
	err := db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(BoltLinksBucket); err != nil {
			return err
		}

		if _, err := tx.CreateBucketIfNotExists(BoltTagsBucket); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &Bolt{
		db,
	}, nil
}
