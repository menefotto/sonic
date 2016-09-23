package operations

import (
	"path"
	"strings"

	"github.com/sonic/lib/confparse"
	"github.com/sonic/lib/store/backends"
	"github.com/sonic/lib/utils/futils"
)

// Removing dependencies is not trivial I haven't really tought about it, I would
// love to get it right once and be done with it tough, I don't like the way
// that RHEL has implemented it, i think I going to study and search a bit the
// internet to find out if someone has already implemented an algorithm to solve
// this problem. I will first think about it, than if I don't come up quickly with
// a nice and simple solution, I will search the internet, but I don't really want
// to reinvent the wheel this time, and work on a really nice algorithms just to
// find out someone has already implemented it.

// temporary implementation quite a shitty one, though it should work

func Remove(pkgname string, p *confparse.IniParser) error {
	testdb, testdir, err := getInstallConf(p)
	if err != nil {
		return err
	}

	db, err := backends.NewBolt(path.Join(testdir, testdb))
	if err != nil {
		return err
	}
	defer db.Close()

	list, err := db.Get([]byte(pkgname))
	if err != nil {
		return err
	}

	err = futils.RemoveList(testdir, strings.Split(string(list), "\n"))
	if err != nil {
		return err
	}

	db.Del([]byte(pkgname))

	return nil
}
