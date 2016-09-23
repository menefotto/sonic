package operations

import (
	"fmt"
	"path"

	"github.com/sonic/lib/confparse"
	"github.com/sonic/lib/dbtostore/pkgtoinfo"
	"github.com/sonic/lib/store"
	"github.com/sonic/lib/store/backends"
)

func Query(pkgname string, parser *confparse.IniParser) error {

	dbname, dblocation, err := getDbConf(parser)
	if err != nil {
		return err
	}

	handle, err := backends.GetBolt(path.Join(dblocation, dbname))
	if err != nil {
		return err
	}

	db := store.New(handle)
	defer db.Close()

	results, err := db.Find(pkgname)
	if err != nil {
		return err
	}

	fmt.Printf("Query result for: %s\t\n", pkgname)
	for name, blob := range results {
		var info pkgtoinfo.PackageInfo
		_, err := blob.Data(&info)
		if err != nil {
		}

		fmt.Printf("%s\n\t%s\n", name, info.Desc)
	}

	return nil
}
