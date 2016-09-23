package pkgneeds

import (
	"fmt"
	"os"
	"testing"

	"github.com/sonic/lib/dbtostore"
)

func TestResolve(t *testing.T) {

	dag, err := dbtostore.Build("../dbtostore/core.db")
	if err != nil {
		t.Logf("Main error: %s\n", err)
		return
	}

	visit := Resolve(dag, "tar")
	for key, _ := range visit {
		fmt.Println(key)
	}
	//t.Logf("Total amount of packages: %v\n", dag.NumVertex())
	os.Remove("sonic.db")
}

func TestGetName(t *testing.T) {
	str1 := "ciao>=3"
	str2 := "ciao>3"
	str3 := "ciao=3"
	s1 := getName(str1)
	s2 := getName(str2)
	s3 := getName(str3)
	if s1 != "ciao" || s2 != "ciao" || s3 != "ciao" {
		t.Fatal("getName not working properly")
	}
}
