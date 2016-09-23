package operations

import (
	"fmt"
	"path"

	"github.com/sonic/lib/confparse"
	"github.com/sonic/lib/dbtostore/pkgtoinfo"
	"github.com/sonic/lib/store"
	"github.com/sonic/lib/store/backends"
)

func getDbConf(p *confparse.IniParser) (string, string, error) {
	dbname, err := p.GetString("database", "dbname")
	if err != nil {
		return "", "", err
	}
	dblocation, err := p.GetString("database", "dblocation")
	if err != nil {
		return "", "", err
	}
	return dbname, dblocation, nil

}

func getDbFromConf(parser *confparse.IniParser) (*backends.Bolt, error) {
	dbname, dblocation, err := getDbConf(parser)
	if err != nil {
		return nil, err
	}

	db, err := backends.GetBolt(path.Join(dblocation, dbname))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getInstallConf(p *confparse.IniParser) (string, string, error) {
	testinstall, err := p.GetString("installation", "installpath")
	if err != nil {
		return "", "", err
	}
	testname, err := p.GetString("installation", "installdb")
	if err != nil {
		return "", "", err
	}
	return testname, testinstall, nil
}

func getDownloadConf(p *confparse.IniParser) (string, string, bool, error) {
	location, err := p.GetString("download", "location")
	if err != nil {
		return "", "", false, err
	}
	pkgcache, err := p.GetString("download", "packagecache")
	if err != nil {
		return "", "", false, err
	}

	silent, err := p.GetBool("download", "silent")
	if err != nil {
		return "", "", false, err
	}

	return location, pkgcache, silent, nil
}

func getRepoUrlConf(p *confparse.IniParser) (string, error) {
	return p.GetString("repos", "baseurl")
}

func getVerifyMethod(pkgname string, p *confparse.IniParser) ([]string, error) {
	info, err := getPkgInfo(pkgname, p)
	if err != nil {
		return nil, err
	}

	return []string{info.Sha256sum, info.Md5sum}, nil
}

func getPkgFileName(pkgname string, p *confparse.IniParser) (string, error) {
	info, err := getPkgInfo(pkgname, p)
	if err == backends.ErrNotFound {
		return "", fmt.Errorf("Target package '%s' not found \n", pkgname)
	}

	return info.Filename, nil
}

func getPkgInfo(pkgname string, p *confparse.IniParser) (*pkgtoinfo.PackageInfo, error) {
	dbname, dblocation, err := getDbConf(p)
	if err != nil {
		return nil, err
	}

	db, err := backends.GetBolt(path.Join(dblocation, dbname))
	if err != nil {
		return nil, err
	}

	localstore := store.New(db)
	defer localstore.Close()

	value, err := localstore.Get(pkgname)
	if err != nil {
		return nil, err
	}

	var info pkgtoinfo.PackageInfo
	_, err = value.Data(&info)
	if err != nil {
		return &info, err
	}
	return &info, nil

}
