package pkgtoinfo

import (
	"fmt"
	"testing"
)

func TestFileToInfo(t *testing.T) {
	info, err := FileToPkgInfo("desc")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("This is filename %v \n", info.Filename)
	fmt.Printf("This is desc %v \n", info.Desc)
	fmt.Printf("This is arch %v \n", info.Arch)
	fmt.Printf("This is csize %s \n", info.Csize)
	fmt.Printf("This is isize %s \n", info.Isize)
	fmt.Printf("This is version %v \n", info.Version)
	fmt.Printf("This is url %v \n", info.Url)
	fmt.Printf("This is license %v \n", info.License)
	fmt.Printf("This is md5sum %s \n", info.Md5sum)
	fmt.Printf("This is packager %v \n", info.Packager)
	//	fmt.Printf("This is pgpsig %v \n", info.Pgpsig)
	fmt.Printf("This are deps %v \n", info.Depends)
}

func TestProvides(t *testing.T) {
	info, err := FileToPkgInfo("depends")
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Provides: ", info.Provides)
}
