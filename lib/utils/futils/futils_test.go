package futils

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/sonic/lib/log"
)

func TestDirList(t *testing.T) {
	directory := "/home/carlo/Work"
	list, err := DirList(directory)
	if err != nil {
		t.Logf("Error %v\n", err)
	}

	for _, dir := range list {
		fmt.Println("Directory :", dir)
	}
}

func TestWriteDirAndFile(t *testing.T) {
	var dmode os.FileMode = 0777
	err := WriteFileOrDir("testdir.txt", []byte(""), os.ModeDir|dmode)
	if err != nil {
		t.Fatal(err)
	}

	var fmode os.FileMode = 0666
	err = WriteFileOrDir("testfile.txt", []byte("ciao"), os.ModeDir|fmode)
	if err != nil {
		t.Fatal(err)
	}

	os.Remove("testdir.txt")
}

func TestCopyFile(t *testing.T) {
	s := "testfile.txt"
	d := "copytest.txt"
	err := CopyFile(s, d)
	if err != nil {
		t.Fatal(err)
	}
	os.Remove(s)
}

func TestMoveFile(t *testing.T) {
	s := "copytest.txt"
	d := "movedtest.txt"
	err := MoveFile(s, d)
	if err != nil {
		t.Fatal(err)
	}
	os.Remove(d)
}

func TestFileMover(t *testing.T) {
	s := "/tmp/"
	d := "/var/cache/pacman/pkg/"
	filename := "tar-1.29-1-x86_64.pkg.tar.xz"

	err := CopyFile(d+filename, s+filename)
	if err != nil {
		t.Fatal(err)
	}

	mover := NewFileMover(log.New("testmover", ""))
	mover.Send(s+filename, d+filename)
	time.Sleep(time.Second * 2)
	mover.Close()
}

func TestFileMoverManyMoreThen32(t *testing.T) {
	s := "/tmp/"
	d := "/var/cache/pacman/pkg/"
	filename := "tar-1.29-1-x86_64.pkg.tar.xz"

	err := CopyFile(d+filename, s+filename)
	if err != nil {
		t.Fatal(err)
	}

	mover := NewFileMover(log.New("testmover", ""))
	for i := 35; i > 0; i-- {
		mover.Send(s+filename, d+filename)
	}
	time.Sleep(time.Second * 0)
	mover.Close()
}

func TestUniquePath(t *testing.T) {
	slice := []string{"/usr/test/", "/usr/test"}
	res := UniquePaths(slice)
	if len(res) != 1 {
		t.Error("shuold be only one path")
	}

}

func TestFileExistYes(t *testing.T) {
	if err := WriteFileOrDir("test.txt", []byte("ciao"), 0666); err != nil {
		t.Fatal(err)
	}
	ok := FileExist("test.txt")
	if !ok {
		t.Fatal("should be ok")
	}
	os.Remove("test.txt")

}

func TestFileExistNoExist(t *testing.T) {
	ok := FileExist("nofile.txt")
	if ok {
		t.Fatal("should not exist")
	}

}

func TestRemove(t *testing.T) {
	err := os.Mkdir("test", 0777)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Mkdir("test/test2", 0777)
	if err != nil {
		t.Fatal(err)
	}

	fname, err := os.Create("test/test2/test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fname.Close()

	list := []string{"test2/"}

	err = RemoveList("test", list)
	if err != nil {
		t.Fatal(err)
	}

	os.RemoveAll("test")
}

func TestRemoveEvalAllCases(t *testing.T) {
	// first case
	err := os.Mkdir("test", 0777)
	if err != nil {
		t.Fatal(err)
	}

	err = RemoveEval("test")
	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Stat("test")
	if err == nil {
		t.Fatal(err)
	}

	// second case
	err = os.Mkdir("test", 0777)
	if err != nil {
		t.Fatal(err)
	}

	fname2, err := os.Create("test/test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fname2.Close()

	err = RemoveEval("test")
	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Stat("test")
	if err != nil {
		t.Fatal(err)
	}
	os.RemoveAll("test")

	// third case
	err = os.Mkdir("test", 0777)
	if err != nil {
		t.Fatal(err)
	}

	fname3, err := os.Create("test/test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fname3.Close()

	err = RemoveEval("test/test.txt")
	if err != nil {
		t.Fatal(err)
	}

	os.Remove("test")

	// 4th case
	err = RemoveEval("testnoexist")
	if err == nil {
		t.Fatal(err)
	}

}
