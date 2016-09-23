package download

import (
	"fmt"
	"os"
	"testing"
)

func TestClietInit(t *testing.T) {
	client := clientInit()
	_, err := client.Get("https://www.google.com")
	if err != nil {
		t.Error("Error: ", err)
	}

}

func TestClietInitWrong(t *testing.T) {
	client := clientInit()
	_, err := client.Get("ht//www.google.com")
	if err == nil {
		t.Error("Should not be nil")
	}

}
func TestSingleDownload(t *testing.T) {
	base := "http://archlinux.polymorf.fr/core/os/x86_64/"
	pkgname := "bash-4.3.046-1-x86_64.pkg.tar.xz"
	err := Single(base, ".", pkgname)
	if err != nil {
		t.Error(err)
	}
	os.Remove(pkgname)
}
func TestSingleDownloadWrong(t *testing.T) {
	base := "http://archlinux.fr/core/os/x86_64/"
	pkgname := "bash-4.3.046-1-x86_64.pkg.tar.xz"
	err := Single(base, ".", pkgname)
	if err == nil {
		t.Error("Should not be nil")
	}
}

func TestMany(t *testing.T) {
	base := "http://archlinux.polymorf.fr/core/os/x86_64/"

	pkgnames := []string{
		"acl-2.2.52-2-x86_64.pkg.tar.xz",
		"bash-4.3.046-1-x86_64.pkg.tar.xz",
	}

	errchan := Many(base, ".", pkgnames)
	for i := 0; i < len(errchan); i++ {
		fmt.Println(errchan[i])
	}

	for _, name := range pkgnames {
		err := os.Remove(name)
		if err != nil {
			//fmt.Println("Err :", err)
		}
	}
}
func TestManyWrong(t *testing.T) {
	base := "http://archlinux.polymorf.fr/core/os/x86_64/"

	pkgnames := []string{
		"acls2-2-x86_64.pkg.tar.xz",
		"bass-4.3.046-1-x86_64.pkg.tar.xz",
	}

	errchan := Many(base, ".", pkgnames)
	for i := 0; i < len(errchan); i++ {
		fmt.Println(errchan[i])
	}

}
func TestDownloadSequential(t *testing.T) {
	base := "http://archlinux.polymorf.fr/core/os/x86_64/"

	pkgnames := []string{
		"acl-2.2.52-2-x86_64.pkg.tar.xz",
		"acl-2.2.52-2-x86_64.pkg.tar.xz",
	}

	for _, pkgname := range pkgnames {
		err := Single(base, ".", pkgname)
		if err != nil {
			t.Error(err)
		}
		os.Remove(pkgname)
	}

}
