package transaction

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/sonic/lib/confparse"
	"github.com/sonic/lib/log"
	"github.com/sonic/lib/operations"
	"github.com/sonic/lib/utils/futils"
)

type Tx struct {
	log        *log.Logger
	dbfilepath string
	statusOk   bool
}

const (
	TxBefore   = ".before_tx"
	TxTmpPath  = "/var/log/sonictx"
	TxNamePath = "/var/log/sonictx/sonic.tx"
	TxExt      = "_tx"
)

func New(dbpath string) *Tx {
	return &Tx{
		log:        log.New("Sonic Tx : ", TxNamePath),
		dbfilepath: dbpath,
		statusOk:   false,
	}
}

func (t *Tx) Begin() error {

	t.log.Log("Begin start \n")

	err := futils.CopyFile(t.dbfilepath, t.dbfilepath+TxBefore)
	if err != nil {
		return err
	}

	t.log.Log("Begin stop: \n")

	return nil
}

// Const ActionI and ActionR are stand for Install action and Remove action
// and are the only two action allowed to be transactional

const (
	ActionI = "I:"
	ActionR = "R:"
)

func (t *Tx) Add(actiontype, filepath string) {
	t.log.Log(actiontype + filepath)
}

func (t *Tx) End() error {

	t.log.Log("End start: \n")

	err := os.Rename(t.dbfilepath, t.dbfilepath+TxExt)
	if err != nil {
		return err
	}

	err = os.Rename(t.dbfilepath+TxBefore, t.dbfilepath)
	if err != nil {
		return err
	}

	err = os.Remove(TxNamePath)
	if err != nil {
		return err
	}

	t.statusOk = true

	t.log.Log("End stop: \n")

	return nil
}

func (t *Tx) RollBack() error {

	if !t.statusOk {
		switch {
		case futils.FileExist(t.dbfilepath + TxBefore):
			if err := os.Remove(t.dbfilepath + TxBefore); err != nil {
				// should log here and perhaps retry
				return nil
			}
			t.rollback()
			fallthrough

		case futils.FileExist(TxNamePath):
			err := os.Remove(TxNamePath)
			if err != nil {
				return nil // should log here, temporary returning nil
			}
		}
	}

	return nil
}

// log line should look like [Sonic Tx : R/I:full file path name]
// should revert the ops in the log file

func (t *Tx) rollback() {
	txf, err := ioutil.ReadFile(TxNamePath)
	if err != nil {
		return
	}

	lines := strings.Split(string(txf), "\n")

	for num := 0; num < len(lines); num++ {
		if num > 1 && len(lines[num]) > 5 {
			t.linerollback(lines[num][LogPrefixLen:])
		}
	}
}

const (
	LogPrefixLen = 20
	LogOpsLen    = 2
)

var (
	InstallAgain func(pkgname string) error = func(pkgname string) error {
		parser, err := confparse.NewParserFromFile("../operations/sonic.conf")
		if err != nil {
			return err
		}

		return operations.Install(pkgname, parser)
	}
)

func (t *Tx) linerollback(fileline string) {
	// dont't really know if I should panic and log or just log or employ error
	// strategy
	switch {
	case fileline[:LogOpsLen] == "R:":
		err := InstallAgain(fileline)
		if err != nil {
			return
		}
	case fileline[:LogOpsLen] == "I:":
		// take into consideration all cases // they should not be
		// ordered and perhaps comming from
		// different files. This is an i.e. kind of non-solution
		err := futils.RemoveEval(fileline[LogOpsLen:])
		if err != nil {
			return
		}
	}
}
