//this package containes files and related utilities
// functions like DirList with returns a slice a names
// with the content in the directory given as parameter
// is it not recurisive, it will not walk all the
// directory tree.

package futils

import (
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/sonic/lib/errors"
)

func DirList(dirname string) ([]string, error) {
	filelist := make([]string, 0)

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return filelist, errors.Wrap(err)()
	}

	for _, file := range files {
		filelist = append(filelist, file.Name())
	}

	return filelist, nil
}

func WriteFileOrDir(name string, data []byte, mode os.FileMode) error {
	if len(data) == 0 {
		return os.MkdirAll(name, mode)
	}

	dirname, _ := path.Split(name)
	if _, err := os.Stat(dirname); err != nil {
		os.MkdirAll(dirname, 0766|mode)
	}

	return ioutil.WriteFile(name, data, mode)
}

//copies file from source to destination the directory in which is run the command
// is going to be added to
func CopyFile(source, destination string) error {
	from, err := os.Open(source)
	if err != nil {
		return errors.Wrap(err)()
	}
	defer from.Close()

	info, err := from.Stat()
	if err != nil {
		return errors.Wrap(err)()
	}

	content := make([]byte, info.Size())
	_, err = from.Read(content)
	if err != nil {
		return errors.Wrap(err)()
	}

	return ioutil.WriteFile(destination, content, info.Mode())
}

func MoveFile(source, destination string) error {
	err := CopyFile(source, destination)
	if err != nil {
		return errors.Wrap(err)()
	}

	err = os.Remove(source)
	if err != nil {
		return errors.Wrap(err)()
	}
	return nil
}

func FileExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil {
		return false
	}

	return true
}

func RemoveList(testdir string, paths []string) error {

	sort.Sort(sort.Reverse(sort.StringSlice(paths)))

	for _, pathentry := range paths {
		//fmt.Println(pathentry)
		if err := RemoveEval(path.Join(testdir, pathentry)); err != nil {
			return err
		}
	}

	return nil
}

func RemoveEval(filepath string) error {
	info, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	if info.IsDir() {
		files, _ := DirList(filepath)

		if len(files) > 0 {
			return nil
		}

		if err := os.RemoveAll(filepath); err != nil {
			return err
		}

		return nil
	}

	if err := os.Remove(filepath); err != nil {
		return err
	}

	return nil
}

func UniquePaths(paths []string) map[string]bool {
	sort.Strings(paths)

	uniq := make(map[string]bool)
	for _, filepath := range paths {
		if len(filepath) > 0 {
			base := strings.Split(filepath, "/")[0]
			uniq[base] = true
		}
	}

	return uniq
}
