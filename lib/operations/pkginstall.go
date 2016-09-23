package operations

import (
	"bytes"
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/sonic/lib/compress/xz"
	"github.com/sonic/lib/confparse"
	"github.com/sonic/lib/log"
	"github.com/sonic/lib/store/backends"
	"github.com/sonic/lib/utils/crypto"
	"github.com/sonic/lib/utils/download"
	"github.com/sonic/lib/utils/futils"
	"github.com/sonic/lib/utils/tar"
)

// install should do dependency resolution first then install all the pkgs
// but first must finish the transaction implementations this is a dummy
// basically

func Install(pkgname string, parser *confparse.IniParser) error {

	// get basic info for installtion
	tempfile, pkgcache, err := setupInstall(pkgname, parser)
	if err != nil {
		return err
	}
	// background cooroutine moves files
	mover := futils.NewFileMover(log.New("mover", ""))

	// download file if not local and perform sanity check
	err = downloadLocal(tempfile, pkgname, parser)
	if err != nil {
		return err
	}

	// fully extract package work sort of like `tar xvf filename`
	archive, err := pkgExtract(tempfile)
	if err != nil {
		return err
	}

	// move the file from temp location to pkgcache
	_, filename := filepath.Split(tempfile)
	mover.Send(tempfile, path.Join(pkgcache, filename))

	// prepare to install !!TODO insert transaction here"
	err = preInstall(pkgname, archive, parser)
	if err != nil {
		return err
	}

	// perform the actual writing to disk
	err = doInstall(archive, parser)
	if err != nil {
		return err
	}

	return nil
}

func pkgExtract(tmpfile string) (map[string]*tar.TarEntry, error) {
	content, err := xz.DecompressFile(tmpfile)
	if err != nil {
		return nil, err
	}

	tararchive, err := tar.Extractor(content)
	if err != nil {
		return nil, err
	}

	return tararchive, nil
}

func setupInstall(pkg string, p *confparse.IniParser) (string, string, error) {

	filename, err := getPkgFileName(pkg, p)
	if err != nil {
		return "", "", err
	}

	tmplocation, pkgcache, _, err := getDownloadConf(p)
	if err != nil {
		return "", "", err
	}

	return path.Join(tmplocation, filename), pkgcache, nil
}

func preInstall(pkgname string, ar map[string]*tar.TarEntry, p *confparse.IniParser) error {
	installdb, basedir, err := getInstallConf(p)
	if err != nil {
		return err
	}

	pathlist := make([]string, 0)
	for file, _ := range ar {

		_, filename := filepath.Split(file)
		if !strings.HasPrefix(filename, ".") {
			pathlist = append(pathlist, file)
		}
	}

	err = trackFiles(pkgname, path.Join(basedir, installdb), pathlist)
	if err != nil {
		return err
	}

	return nil
}

func doInstall(ar map[string]*tar.TarEntry, p *confparse.IniParser) error {
	_, basedir, err := getInstallConf(p)
	if err != nil {
		return err
	}

	for fpath, entry := range ar {
		file := path.Join(basedir, fpath)
		// temp solution only for testing

		_, filename := filepath.Split(file)
		if !strings.HasPrefix(filename, ".") {
			err := futils.WriteFileOrDir(file, entry.Data, entry.Mode)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func downloadLocal(tmpfile, pkgname string, parser *confparse.IniParser) error {

	tmplocation, filename := filepath.Split(tmpfile)

	if !isLocal(filename) {
		url, err := parser.GetString("repos", "baseurl")
		if err != nil {
			return err
		}

		err = download.Single(url, tmplocation, filename)
		if err != nil {
			return err
		}
	} else {
		err := futils.CopyFile(filename, tmpfile)
		if err != nil {
			return err
		}
	}

	err := verifyPackage(tmpfile, pkgname, parser)
	if err != nil {
		return err
	}

	return nil
}

func isLocal(pkgname string) bool {
	_, err := os.Stat(pkgname)
	if err != nil {
		return false
	}
	return true
}

func verifyPackage(tmpfile, pkgname string, parser *confparse.IniParser) error {
	cryptosum, err := getVerifyMethod(pkgname, parser)
	if err != nil {
		return err
	}

	switch {
	case len(cryptosum[0]) > 0:
		_, err := crypto.VerifySha256(cryptosum[0], tmpfile)
		if err != nil {
			return err
		}
	case len(cryptosum[1]) > 0:
		_, err := crypto.VerifyMd5(cryptosum[1], tmpfile)
		if err != nil {
			return err
		}
	default:
		return errors.New("not verfication method found in db")
	}

	return nil
}

func trackFiles(pkgname, dbname string, paths []string) error {
	db, err := backends.NewBolt(dbname)
	if err != nil {
		return err
	}
	defer db.Close()

	// case the file is being reinstalled the precendent version will be
	// deleted from the database holding the installed packages
	_, err = db.Get([]byte(pkgname))
	if err != backends.ErrNotFound {
		db.Del([]byte(pkgname))
	}

	buff := bytes.NewBuffer(make([]byte, 0))
	for _, value := range paths {
		buff.WriteString(value + "\n")
	}

	err = db.Put([]byte(pkgname), buff.Bytes())
	if err != nil {
		return err
	}

	return nil
}
