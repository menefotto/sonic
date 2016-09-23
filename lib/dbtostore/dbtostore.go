// This package expose one functio. Build given an arch database returns
// a dependecy graph ( directed a-cyclic weighted graph ) and returns, really
// important : the current implementation current implementation demands that
// is the caller to close the underling database once done with the graph.
// Like this :
// c := g.BackEnd()
// defer c.Close()
// I know is not a beautiful solution but what can I say whather is "good enough"
// for the time being, then will see about it.

package dbtostore

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"strings"

	"github.com/sonic/lib/dbtostore/pkgtoinfo"
	"github.com/sonic/lib/errors"
	"github.com/sonic/lib/store"
	"github.com/sonic/lib/store/backends"
)

func Build(archdbname string) (*store.Store, error) {
	f, err := os.Open(archdbname)
	if err != nil {
		return nil, errors.Wrap(err)()
	}
	defer f.Close()

	gizip, err := gzip.NewReader(f)
	if err != nil {
		return nil, errors.Wrap(err)()
	}
	defer gizip.Close()

	archive := tar.NewReader(gizip)
	if err != nil {
		return nil, errors.Wrap(err)()
	}

	return archiveToStore(archive)
}

var PkgDbName string = "sonic.db"

func archiveToStore(a *tar.Reader) (*store.Store, error) {
	dbCloser := func(db *backends.Bolt) {
		if db != nil {
			db.Close()
		}
	}

	db, err := backends.NewBolt(PkgDbName)
	if err != nil {
		defer dbCloser(db)
		return nil, errors.Wrap(err)()
	}

	g := store.New(db)

	infomap := make(map[string][]byte, 0)
	var dirname string

	for {
		header, err := a.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			defer dbCloser(db)
			return nil, errors.Wrap(err)()
		}

		fillMap(infomap, a, header, &dirname)
	}

	for _, v := range infomap {
		err := addToStore(v, g)
		if err != nil {
			defer dbCloser(db)
			return nil, errors.Wrap(err)()
		}
	}

	return g, nil
}

const (
	FILETYPE string = "0"
	DIRTYPE  string = "5"
)

func fillMap(inf map[string][]byte, a *tar.Reader, h *tar.Header, dir *string) {

	if string(h.Typeflag) == DIRTYPE {
		*dir = h.Name
	}

	if string(h.Typeflag) == FILETYPE {
		if strings.Contains(string(h.Name), *dir) {

			data, ok := readEntry(a, h)
			if ok {
				inf[*dir] = append(inf[*dir], data...)
			}
		}

	}

}

func readEntry(a *tar.Reader, h *tar.Header) ([]byte, bool) {
	data := make([]byte, h.Size)

	n, err := a.Read(data)
	if err == io.EOF || err != nil {
		return nil, false
	}
	if n == 0 {
		return nil, false
	}

	return data, true

}

func addToStore(d []byte, dag *store.Store) error {
	info := pkgtoinfo.StringToPkgInfo(string(d))

	err := dag.Add(info.Name, &info)
	if err != nil {
		return errors.Wrap(err)()
	}

	return nil
}
