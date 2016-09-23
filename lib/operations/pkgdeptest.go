package operations

import (
	"fmt"
	"sort"

	"github.com/sonic/lib/confparse"
	"github.com/sonic/lib/pkgneeds"
	"github.com/sonic/lib/store"
)

func DepTest(pkgname string, parser *confparse.IniParser) error {
	sortedpkgs, err := DepSolve(pkgname, parser)
	if err != nil {
		return err
	}

	fmt.Printf("Dependency resolution result for package: %s\n", pkgname)
	for _, value := range sortedpkgs {
		fmt.Printf("\t%s\n", value)
	}

	return nil
}

func DepSolve(pkgname string, parser *confparse.IniParser) ([]string, error) {
	db, err := getDbFromConf(parser)
	if err != nil {
		return nil, err
	}

	newgraph := store.New(db)
	defer newgraph.Close()

	return sortPackages(pkgneeds.Resolve(newgraph, pkgname)), nil
}

func sortPackages(visited map[string]bool) []string {

	sorted := make([]string, 0)
	for key, _ := range visited {
		sorted = append(sorted, key)
	}
	sort.Strings(sorted)

	return sorted
}
