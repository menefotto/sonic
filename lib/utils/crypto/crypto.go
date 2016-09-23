package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sonic/lib/errors"
)

func FileSha256Sum(filename string) (string, error) {

	f, err := os.Open(filename)
	if err != nil {
		return "", errors.Wrap(err)()
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return "", errors.Wrap(err)()

	}

	return fmt.Sprintf("%x", sha256.Sum256(content)), nil

}

func VerifySha256(sha, filename string) (bool, error) {
	newsha, err := FileSha256Sum(filename)
	if err != nil {
		return false, errors.Wrap(err)()
	}

	return newsha == sha, nil

}

func FileMd5Sum(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", errors.Wrap(err)()
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return "", errors.Wrap(err)()
	}

	return fmt.Sprintf("%x", md5.Sum(content)), nil

}

func VerifyMd5(md5, filename string) (bool, error) {
	newmd5, err := FileMd5Sum(filename)
	if err != nil {
		return false, errors.Wrap(err)()
	}

	return newmd5 == md5, nil
}
