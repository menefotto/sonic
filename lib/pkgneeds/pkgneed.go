// This package provides a single function called
// Resolve given a graph and a starting edge
// resolves all the dependecies and they are put
// inside the visited map

package pkgneeds

import (
	"strings"

	"github.com/sonic/lib/dbtostore/pkgtoinfo"
	"github.com/sonic/lib/errors"
	"github.com/sonic/lib/store"
)

func Resolve(dag *store.Store, name string) map[string]bool {
	v := make(map[string]bool, 0)
	s := make(map[string]bool, 0)
	resolver(dag, name, v, s)

	return v
}

func resolver(d *store.Store, e string, v map[string]bool, s map[string]bool) {
	s[getName(e)] = true

	edgelist, err := getEdgeList(d, getName(e))
	if err != nil {
		return
	}

	for idx := 0; idx < len(edgelist); idx++ {
		edge := edgelist[idx]

		if _, ok := v[edge]; !ok {
			if _, ok := s[edge]; ok && (idx+1) < len(edgelist) {
				edgenext := edgelist[idx+1]

				resolver(d, edgenext, v, s)
			} else {
				resolver(d, edge, v, s)
			}
		}
	}

	v[getName(e)] = true

}
func getEdgeList(dag *store.Store, edgename string) ([]string, error) {
	edge, err := dag.Get(edgename)
	if err != nil {
		return nil, errors.Wrap(err)()
	}

	info := &pkgtoinfo.PackageInfo{}
	_, err = edge.Data(&info)
	if err != nil {
		return nil, err
	}

	return info.Depends, nil
}

func getName(name string) string {
	var namesplitted string

	if strings.Contains(name, ">") && strings.Contains(name, "=") {
		namesplitted = strings.Split(name, ">")[0]
		return namesplitted
	}
	if strings.Contains(name, "=") {
		namesplitted = strings.Split(name, "=")[0]
		return namesplitted
	}
	if strings.Contains(name, ">") {
		namesplitted = strings.Split(name, ">")[0]
		return namesplitted
	}

	return name

}
