package main

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/sonic/lib/dbtostore"
	"github.com/sonic/lib/utils/futils"
)

func main() {
	const dbpath string = "/var/lib/pacman/sync"

	dbs, err := futils.DirList(dbpath)
	if err != nil {
		fmt.Println(err)
		return
	}

	sort.Strings(dbs)

	fmt.Printf("Starting to covert %+v into sonic.db: \n", dbs)
	for _, dbname := range dbs {
		store, err := dbtostore.Build(filepath.Join(dbpath, dbname))
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		store.Close()

		fmt.Printf("Builded %s into sonic.db\n", dbname)
	}

}
