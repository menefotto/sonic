package store

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"os"
	"testing"

	"github.com/sonic/lib/dbtostore/pkgtoinfo"
	"github.com/sonic/lib/store/backends"
)

func TestStore(t *testing.T) {
	var db backends.Bolt
	err := db.Open("test.db")
	if err != nil {
		t.Fatal(err)
	}

	g := New(&db)

	_ = g.Add("stovari", []byte("miao"))
	_ = g.Add("carlo", []byte("ciao"))

	_, err = g.Get("stovari")
	if err != nil {
		t.Fatal(err)
	}
	_, err = g.Get("storie")
	if err == nil {
		t.Fatal(err)
	}

	_ = g.BackEnd()
	g.Close()
	os.Remove("test.db")
}

func TestStoreAddDelGet(t *testing.T) {
	var db backends.Bolt
	err := db.Open("test.db")
	if err != nil {
		t.Error(err)
	}

	g := New(&db)
	defer g.Close()

	_ = g.Add("carlo", []byte("ciao"))

	res, err := g.Find("c*")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", res)

	v, err := g.Get("carlo")
	if err != nil {
		t.Errorf("Error in Get %v\n", err)
	}

	var ori []byte
	_, err = v.Data(&ori)
	if err != nil {
		t.Fatal(err)
	}

	if string(ori) != "ciao" {
		t.Fatal("boh dont match")
	}

	g.Del("carlo")
	if db.Len() != 0 {
		t.Errorf("Failed to delete the key/value key num is %d\n", db.Len())
	}

	err = PrettyPrint(db)
	if err == nil {
		t.Error("should have not printed")
	}
	var info pkgtoinfo.PackageInfo
	var info3 pkgtoinfo.PackageInfo

	info.Filename = "tar-193.4.xz"
	_ = g.Add("tar", info)
	fname, err := g.Get("tar")
	if err != nil {
		t.Fatal(err)
	}
	_, err = fname.Data(&info3)
	if err != nil {
		t.Fatal(err)
	}
	if info.Filename != info3.Filename {
		t.Fatal("missmatch shit")
	}
	os.Remove("test.db")
}

func TestStoreMemory(t *testing.T) {
	db := backends.NewMap()
	g := New(db)
	defer g.Close()
	_ = g.Add("carlo", []byte("locci"))

	val, err := g.Get("carlo")
	if err != nil {
		t.Errorf("Error in Get %v\n", err)
	}

	var s []byte
	_, err = val.Data(&s)
	if err != nil {
		t.Fatal(err)
	}

	if string(s) != "locci" {
		t.Errorf("they should be equal %s, %s\n", s, "locci")
	}

	err = MarshallToDisk(db)
	if err != nil {
		t.Fatal(err)
	}

	dbnew := backends.NewMap()
	err = UnmarshalFromDisk(dbnew)
	if err != nil {
		t.Fatal(err)
	}
	err = UnmarshalFromDisk("test")
	if err == nil {
		t.Fatal("shoul not be possible to unmarshall a string")
	}
	err = MarshallToDisk("test")
	if err == nil {
		t.Fatal("shoul not be possible to unmarshall a string")
	}
	err = PrettyPrint(db)
	if err != nil {
		t.Error(err)
	}

}

func TestAny(t *testing.T) {
	var empty string = "XXXXXXX"
	gz := gzip.NewWriter(bytes.NewBuffer([]byte(empty)))

	a := NewAny(nil)
	a.Data(gz)
}
