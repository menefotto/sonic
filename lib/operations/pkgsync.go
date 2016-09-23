package operations

import (
	"github.com/sonic/lib/confparse"
	"github.com/sonic/lib/utils/download"
)

func Sync(value bool, parser *confparse.IniParser) error {

	dbname, dblocation, err := getDbConf(parser)
	if err != nil {
		return err
	}

	baseurl, err := getRepoUrlConf(parser)
	if err != nil {
		return err
	}

	err = download.Single(baseurl, dblocation, dbname)
	if err != nil {
		return err
	}

	return nil
}
