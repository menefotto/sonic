// package sonicdb implementes the backends for
// the graph data type. It is defined by an interface
// containing three methods Get Put and Delete
// anything that satisfy this interface can be used
// and has to be used during the graph datatype initialization
// The Mem type exposed is an in memory backed db implementation
// and it's used for testing any database has to implement this methods
// to satisfy the interface and can be used as a backend
// One note on the Del method in case of failure of that
// is key not found is should not return anything or do
// anything much like the Mem type in the go language itself.
// Bolt provides an transactional
// backend implemented using boltdb provides 5 methods
// Get, Put, Del, Open and Close, three of these are required
// by the db interface

package backends

import (
	"bytes"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/sonic/lib/errors"
)

type DB interface {
	Get(value []byte) ([]byte, error)
	Put(key []byte, value []byte) error
	Del(key []byte)
	Query(key []byte, t string) (map[string][]byte, error)
	Close() error
}

var ErrNotFound error = fmt.Errorf("key not found\n")

type Bolt struct {
	Db         *bolt.DB
	bucketname string
}

func NewBolt(dbname string) (*Bolt, error) {
	return getBolt(dbname)
}

func GetBolt(dbname string) (*Bolt, error) {
	return getBolt(dbname)
}

func getBolt(dbname string) (*Bolt, error) {
	var db Bolt
	err := db.Open(dbname)
	if err != nil {
		return nil, err
	}
	return &db, nil
}

func (b *Bolt) Open(name string) error {
	var err error
	b.Db, err = bolt.Open(name, 0600, nil)
	if err != nil {
		return errors.Wrap(err)()
	}

	b.bucketname = "all"

	_ = b.Db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(b.bucketname))
		if err != nil {
			return errors.Wrap(err)()
		}
		return nil
	})

	return nil
}

func (b *Bolt) Close() error {
	if b.Db != nil {
		return b.Db.Close()
	}
	return nil
}

func (b *Bolt) Get(value []byte) (result []byte, err error) {
	// err assigned directly to err named return value
	err = b.Db.View(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte(b.bucketname))
		// not checking for exisistance since it always exist
		result = buck.Get(value)
		// value assigned directly to named return value
		if result == nil {
			return errors.Wrap(ErrNotFound)()
		}
		return nil
	})
	return
}

func (b *Bolt) Put(key []byte, value []byte) error {
	err := b.Db.Update(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte(b.bucketname))
		// not checking since it alway exist
		err := buck.Put(key, value)
		if err != nil {
			return errors.Wrap(err)()
		}
		return nil
	})
	return err

}

func (b *Bolt) Del(key []byte) {
	_ = b.Db.Update(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte(b.bucketname))
		// not checking for error since it always exist
		buck.Delete(key)
		// not checking for error since it always exist
		// and in case it doesn't we why raise an error
		// follow the Mem delete api
		return nil
	})
	return
}

func (b *Bolt) Len() (value int) {
	// err is not taken into consideration since it always nil
	_ = b.Db.View(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte(b.bucketname))
		// not checking for exisistance since it always exist
		stat := buck.Stats()
		// value assigned directly to named return value
		value = stat.KeyN
		return nil
	})
	return
}

func (b *Bolt) Query(name []byte, t string) (result map[string][]byte, err error) {
	err = b.Db.View(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte(b.bucketname))
		// got the cursor now iterates over the k, values
		result = make(map[string][]byte, 0)
		// iterate over the elements using hasPrefix
		var Search func(s, name []byte) bool

		switch {
		case t == "p":
			Search = bytes.HasPrefix
			name = name[:len(name)-1]
		case t == "s":
			Search = bytes.HasSuffix
			name = name[1:]
		}

		buck.ForEach(func(k, v []byte) error {
			if Search(k, name) {
				result[string(k)] = v
			}
			return nil
		})

		return nil
	})

	return
}

type Mem struct {
	Db map[string]string
}

func NewMap() *Mem {
	return &Mem{Db: make(map[string]string, 0)}
}

func (m *Mem) Get(key []byte) ([]byte, error) {
	value, ok := m.Db[string(key)]
	if ok {
		return []byte(value), nil
	}
	return nil, ErrNotFound
}

func (m *Mem) Put(key []byte, value []byte) error {
	m.Db[string(key)] = string(value)
	// always return nil it cannot fail if it fails Mem implementaion
	// is going to panic no point double check again
	return nil

}

func (m *Mem) Del(key []byte) {
	delete(m.Db, string(key))
}

func (m *Mem) Close() error {

	// does nothing here to satisfy the DB interface
	return nil
}

func (m *Mem) Query(key []byte, str string) (map[string][]byte, error) {
	return nil, errors.New("not implemented")
}
