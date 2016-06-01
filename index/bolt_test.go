package index_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/nproc/kb/index"
	"github.com/nproc/kb/link"
	"github.com/nproc/kb/url"
	"github.com/nproc/kb/user"
	"github.com/stretchr/testify/assert"
)

var db *bolt.DB

func TestMain(m *testing.M) {
	file, err := ioutil.TempFile("", "")

	if err != nil {
		return
	}

	bdb, err := bolt.Open(file.Name(), 0600, nil)

	if err != nil {
		return
	}

	db = bdb

	exitCode := m.Run()

	os.Remove(file.Name())
	bdb.Close()

	os.Exit(exitCode)
}

func TestNewBoltIndex(t *testing.T) {
	index, err := index.NewBoltIndex(db)

	assert.NotNil(t, index)
	assert.Nil(t, err)
}

func TestBoltAdd(t *testing.T) {
	link := link.New(
		url.NewOrNil("https://google.com"),
		user.New("Andrey", "andreykvital"),
		[]string{"#google", "#go", "#search"},
	)

	idx, err := index.NewBoltIndex(db)

	assert.Nil(t, err)
	assert.NotNil(t, idx)

	err = idx.Add(link)

	assert.Nil(t, err)

	db.View(func(tx *bolt.Tx) error {
		serialized := tx.Bucket(index.BoltLinksBucket).Get([]byte(link.ID().String()))

		assert.NotNil(t, serialized)
		assert.NotEmpty(t, serialized)

		tx.Bucket(index.BoltTagsBucket).ForEach(func(key, value []byte) error {
			var (
				tag    = string(key)
				linkID = string(value)
			)

			assert.Equal(t, link.ID().String(), linkID)
			assert.Contains(t, link.Tags(), tag)

			return nil
		})

		return nil
	})
}

func TestBoltAddKeepsOnlyOneLinkReferencePerTag(t *testing.T) {
}

func TestBoltAddOverridePreviouslyIndexedURL(t *testing.T) {
}
