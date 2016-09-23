package dbtostore

import (
	"fmt"
	"testing"

	"github.com/sonic/lib/dbtostore/pkgtoinfo"
	"github.com/sonic/lib/utils/futils"
)

func TestPrepare(t *testing.T) {
	err := futils.CopyFile("/var/lib/pacman/sync/core.db", "core.db")
	if err != nil {
		t.Error(err)
	}

}

func TestDbToStore(t *testing.T) {
	g, err := Build("core.db")
	if err != nil {
		t.Error(err)
		return
	}
	defer g.Close()

	v, err := g.Get("gcc")
	if err != nil {
		t.Error(err)
		return
	}

	info := &pkgtoinfo.PackageInfo{}
	_, err = v.Data(&info)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Info name is %v\n", info.Name)
	fmt.Printf("Info dependes is %v\n", info.Depends)
	fmt.Printf("Info makedepemds is  %v\n", info.Makedepends)
	//fmt.Printf("Info pgpsig is %v\n", info.Pgpsig)

	//os.Remove("packages.db")
}

func TestArchDbNonExistDb(t *testing.T) {
	g, err := Build("rubbish.db")
	if err == nil {
		t.Error(err)
		return
	}

	if g != nil {
		t.Error("store should be nil")
	}

}

func TestArchDbMalformedGzipDb(t *testing.T) {
	g, err := Build("core_failgz.db")
	if err == nil {
		t.Error(err)
		return
	}

	if g != nil {
		t.Error("store should be nil")
	}

}

func TestArchDbNotGzipDb(t *testing.T) {
	g, err := Build("core_notgz.db")
	if err == nil {
		t.Error(err)
		return
	}

	if g != nil {
		t.Error("store should be nil")
	}
}

func TestArchAddToStore(t *testing.T) {
	g, err := Build("core.db")
	if err != nil {
		t.Error(err)
		return
	}

	err = addToStore([]byte("ciao"), g)
	if err == nil {
		t.Fatal(err)
	}

}

func TestNotEnoughPermDbLocation(t *testing.T) {
	PkgDbName = "/var/cache/fake.db"
	_, err := Build("core.db")
	if err == nil {
		t.Error(err)
		return
	}
}
