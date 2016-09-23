package transaction

import (
	"os"
	"testing"

	"github.com/sonic/lib/utils/futils"
)

func TestTxOk(t *testing.T) {
	path := "/home/wind85/Work/go/src/github.com/sonic/tests/"
	installdb := "install.db"
	mockfile := "test.db"

	tx := New(path + installdb)

	err := tx.Begin()
	if err != nil {
		t.Fatal(err)
	}

	err = futils.CopyFile(path+"sonic.db", "/tmp/"+mockfile)
	if err != nil {
		t.Fatal(err)
	}

	tx.Add(ActionI, "/tmp/"+mockfile)

	err = tx.End()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err := tx.RollBack()
		if err != nil {
			panic(err)
		}
	}()

	//os.Remove("/tmp/" + mockfile)
}

func TestLinerollback(t *testing.T) {
	path := "/home/wind85/Work/go/src/github.com/sonic/tests/"
	installdb := "install.db"

	_, err := os.Stat("/tmp/test.db")
	if err != nil {
		t.Fatal(err)
	}

	tx := New(path + installdb)
	line := "Sonic Tx : 11:07:49 I:/tmp/test.db"
	tx.linerollback(line[LogPrefixLen:])

	_, err = os.Stat("/tmp/test.db")
	if err == nil {
		t.Fatal(err)
	}
}

func TestTxLeaveTxfile(t *testing.T) {
	path := "/home/wind85/Work/go/src/github.com/sonic/tests/"
	installdb := "install.db"
	mockfile := "test.db"

	tx := New(path + installdb)

	err := tx.Begin()
	if err != nil {
		t.Fatal(err)
	}

	err = futils.CopyFile(path+"sonic.db", "/tmp/"+mockfile)
	if err != nil {
		t.Fatal(err)
	}

	tx.Add(ActionI, "/tmp/"+mockfile)
}

func TestRollBack(t *testing.T) {
	path := "/home/wind85/Work/go/src/github.com/sonic/tests/"
	installdb := "install.db"
	mockfile := "test.db"

	err := futils.CopyFile(path+"sonic.db", "/tmp/"+mockfile)
	if err != nil {
		t.Fatal(err)
	}

	tx := New(path + installdb)
	tx.rollback()

	_, err = os.Stat("/tmp/test.db")
	if err == nil {
		t.Fatal(err)
	}

}

func TestTxNotOk(t *testing.T) {
	path := "/home/wind85/Work/go/src/github.com/sonic/tests/"
	installdb := "install.db"
	mockfile := "test.db"

	tx := New(path + installdb)

	err := tx.Begin()
	if err != nil {
		t.Fatal(err)
	}

	err = futils.CopyFile(path+"sonic.db", "/tmp/"+mockfile)
	if err != nil {
		t.Fatal(err)
	}

	tx.Add(ActionI, "/tmp/"+mockfile)

	//instead of paniching

	defer func() {
		err := tx.RollBack()
		if err != nil {
			panic(err)
		}
	}()

}

func TestTxNotOkWorked(t *testing.T) {
	_, err := os.Stat("/tmp/test.db")
	if err == nil {
		t.Fatal("Transaction failure\n")
	}
}

func TestInstallAgain(t *testing.T) {
	err := InstallAgain("tar")
	if err != nil {
		t.Fatal(err)
	}

	os.RemoveAll("../../tests/usr")
}
