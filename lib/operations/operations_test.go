package operations

import (
	"log"
	"os"
	"path"
	"testing"

	"github.com/sonic/lib/confparse"
	"github.com/sonic/lib/utils/futils"
)

func TestPrepare(t *testing.T) {
	spath := "/home/wind85/Work/go/src/github.com/sonic/lib/dbtostore"
	source := path.Join(spath, "sonic.db")
	err := futils.CopyFile(source, path.Join("/home/wind85/Work/go/src/github.com/sonic/tests", "sonic.db"))
	if err != nil {
		t.Fatal(err)
	}
}

type global struct {
	P *confparse.IniParser
	E error
}

func NewGlobal() *global {
	parse, err := confparse.NewParserFromFile("sonic.conf")
	if err != nil {
		log.Fatal(err)
	}
	return &global{P: parse, E: err}
}

var g *global = NewGlobal()

func TestOpDepTest(t *testing.T) {
	pkgname := "tar"
	err := DepTest(pkgname, g.P)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpPkgQuerySuffix(t *testing.T) {
	search := "*ar"
	err := Query(search, g.P)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpPkgQueryPrefix(t *testing.T) {
	search := "libn*"
	err := Query(search, g.P)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpPkgInstall(t *testing.T) {
	err := Install("tar", g.P)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpPkgRemove(t *testing.T) {
	err := Remove("tar", g.P)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSync(t *testing.T) {
	err := Sync(true, g.P)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClean(t *testing.T) {
	os.Remove("../../testdir/install.db")
	os.Remove("packages.db")
}

func TestGetIfNotLocal(t *testing.T) {
	p, err := confparse.NewParserFromFile("sonic.conf")
	if err != nil {
		t.Fatal(err)
	}

	downloadLocal("", "", p)
}
