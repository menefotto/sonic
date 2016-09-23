// The Store package provide a generelized data structure ideal for manupulating
// Store in a simple fashion, by providing a backend to it, let's you choose if
// use want to have persistency of a ram based approach.
// To obtain generalization the edges are a simple interface the implements the
// Data() []bytes method, down here there is an example Any, in practise an edge
// must serialize itself and take care of it thanks to the go standard library
// this is quite an easy job.

package store

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/sonic/lib/compress/snappy"
	"github.com/sonic/lib/errors"
	"github.com/sonic/lib/store/backends"
)

// defines the StoreItem interface
type StoreItem interface {
	Data(v interface{}) ([]byte, error)
}

//implements an StoreItem interface
type Any struct {
	Buffer []byte
	data   interface{}
}

func NewAny(v interface{}) *Any {
	return &Any{
		Buffer: nil,
		data:   v,
	}
}

// Data is the only method in the store item iterface and it is implemented like
// a toogle meaning that depending on the internal state of Any the struct that
// implements the Store Item interface has different behaviors.
// I.E:
// there are two different types of way to create an Any either with a call to NewAny(v)
// or the struct literal way &Any{Buffer:data,data:nil}
// The call on data checks whether or not the internal buffer is nil or valid
// if it's nil the value interface type gets parsed into json and put into the internal
// buffer.
// The call on data if the internal buffer is not nul gets unmasheled back to the data
// structure passed as interface value and the buffer gets set to nil again.

func (a *Any) Data(v interface{}) ([]byte, error) {
	var err error

	if a.Buffer == nil {
		a.Buffer, err = json.Marshal(&a.data)
		if err != nil {
			return nil, errors.Wrap(err)()
		}
		return a.Buffer, nil
	}

	err = json.Unmarshal(a.Buffer, v)
	if err != nil {
		return nil, errors.Wrap(err)()
	}
	a.Buffer = nil

	return nil, nil
}

//Store implementation

var ErrNotSupported error = fmt.Errorf("back end not supported")

type Store struct {
	db backends.DB
}

func New(g backends.DB) *Store {
	return &Store{db: g}
}

func (g *Store) Add(key string, value interface{}) error {
	data, err := NewAny(value).Data(nil)
	if err != nil {
		return err
	}

	err = g.db.Put([]byte(key), snappy.Compress(data))
	if err != nil {
		return err
	}

	return nil
}

func (g *Store) Get(key string) (StoreItem, error) {
	val, err := g.db.Get([]byte(key))
	if err != nil {
		return nil, err
	}

	data, err := snappy.Decompress(val)
	if err != nil {
		return nil, err
	}

	return &Any{Buffer: data}, nil
}

func (g *Store) Find(query string) (map[string]StoreItem, error) {
	var queryt string

	switch idx := strings.Index(query, "*"); {
	case idx == 0:
		queryt = "s"
	case idx == len(query)-1:
		queryt = "p"
	default:
		queryt = "p"
	}

	results, err := g.db.Query([]byte(query), queryt)
	if err != nil {
		return nil, err
	}

	anymap := make(map[string]StoreItem, 0)
	for k, val := range results {
		data, err := snappy.Decompress(val)
		if err != nil {
			return nil, err
		}
		results[k] = nil

		anymap[k] = &Any{Buffer: data}
	}

	return anymap, nil
}

func (g *Store) Del(key string) {
	g.db.Del([]byte(key))

}

func (g *Store) Close() {
	switch g.db.(type) {
	case *backends.Bolt:
		bolt := g.db.(*backends.Bolt)
		bolt.Close()
	default:
		return
	}
}

func (g *Store) BackEnd() backends.DB {
	return g.db
}

//helper functions for debugging memory backend
// do not implement better errors

func MarshallToDisk(g interface{}) error {
	switch g.(type) {
	case *backends.Mem:
		b, err := json.Marshal(g)
		if err != nil {
			return err
		}

		f, err := os.Create("store.dump")
		defer f.Close()
		if err != nil {
			return err
		}

		n, err := f.Write(b)
		if n != len(b) || err != nil {
			return err
		}
	default:
		return ErrNotSupported
	}
	return nil
}

func UnmarshalFromDisk(g interface{}) error {
	switch g.(type) {
	case *backends.Mem:
		f, err := os.Open("store.dump")
		defer f.Close()
		if err != nil {
			return err
		}

		stat, err := f.Stat()
		if err != nil {
			return err
		}

		data := make([]byte, stat.Size())
		count, err := f.Read(data)
		if count != int(stat.Size()) || err != nil {
			return fmt.Errorf("missmatched size\n")
		}

		err = json.Unmarshal(data, g)
		if err != nil {
			return err
		}
	default:
		return ErrNotSupported
	}

	return nil
}

func PrettyPrint(g interface{}) error {
	switch t := g.(type) {
	case *backends.Mem:
		for key, value := range t.Db {
			fmt.Println("-----------------------------------------")
			fmt.Println("Key is: ", key)
			for _, val := range value {
				fmt.Printf("\t\t Value is: %v\n", val)
			}
			fmt.Println("-----------------------------------------")
		}
		return nil
	default:
		return ErrNotSupported

	}
}
