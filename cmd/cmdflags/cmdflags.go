//package cmdline provides the CommandExecutor() function who does
//parse the command line arguments and forward them to the right function
// the Operations availables are defined on the Operations var, the struct
// itself is composed by a series of functions with the Function type signature
// that is func(v string) error, each of the function must be defined by the
// caller in this case the main function.

package cmdflags

import (
	"errors"
	"fmt"
	"os"

	"github.com/sonic/lib/confparse"
	flag "github.com/spf13/pflag"
)

type Function func(pkg string, parser *confparse.IniParser) error
type FunctionBool func(value bool, parser *confparse.IniParser) error

var Operations struct {
	InstallPkgs Function
	RemovePkgs  Function
	QueryPkgs   Function
	DeptestPkgs Function
	SyncDbs     FunctionBool
}

var (
	ErrNotRoot = errors.New("You must be root to execute this command!\n")
	ErrZeroOps = errors.New("error: No operation specified (type -h for help)\n")
)

func CommandExecutor(p *confparse.IniParser) error {
	var err error
	flag.Parse()

	if flag.NFlag() == 0 {
		return ErrZeroOps
	}

	switch {
	case opts.Sync:
		err = Operations.SyncDbs(opts.Sync, p)
		if flag.NFlag() == 1 {
			break
		}

		fallthrough
	case hasCmd(opts.Install):
		if !isRoot() {
			return ErrNotRoot
		}
		err = Operations.InstallPkgs(opts.Install, p)
	case hasCmd(opts.Remove):
		if !isRoot() {
			return ErrNotRoot
		}
		err = Operations.RemovePkgs(opts.Remove, p)
	case hasCmd(opts.Query):
		err = Operations.QueryPkgs(opts.Query, p)
	case hasCmd(opts.Deptest):
		err = Operations.DeptestPkgs(opts.Deptest, p)
	case opts.Version:
		fmt.Println("Sonic Version 0.1")
	}

	return err
}

func isRoot() bool {
	if uid := os.Getuid(); uid != 0 {
		return false
	}
	return true
}

func hasCmd(v string) bool {
	if len(v) >= 1 {
		return true
	}
	return false
}

// type cmd hold the command value type in form of string satifies the flag-value
// interface defined by the go stdlib package flag

type cmd struct {
	pkgs string
}

func (c *cmd) String() string {
	return c.pkgs
}

func (c *cmd) Set(value string) error {
	c.pkgs = value
	return nil
}

func (c *cmd) Type() string {
	return "string"
}

// opts contains a the set of ops supported by sonic, in particular this structure
// holds all the results arguments once parsed from the command line.
var opts struct {
	Sync    bool
	Install string
	Remove  string
	Query   string
	Deptest string
	Version bool
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, Help)
	}

	flag.StringVarP(&opts.Install, "install", "i", "", install)
	flag.StringVarP(&opts.Remove, "remove", "r", "", remove)
	flag.StringVarP(&opts.Query, "query", "q", "", query)
	flag.StringVarP(&opts.Deptest, "deptest", "d", "", deptest)
	flag.BoolVarP(&opts.Sync, "sync", "s", false, sync)
	flag.BoolVarP(&opts.Version, "version", "v", false, version)
}

const (
	usage      = "usage:  sonic <operation> [...]"
	operations = "operations:"
	install    = "sonic {-i --install}  [package(s)]"
	remove     = "sonic {-r --remove}   [package(s)]"
	query      = "sonic {-q --query}    [package(s)]"
	deptest    = "sonic {-d --deptest}  [package(s)]"
	sync       = "sonic {-s --sync}"
	version    = "sonic {-v --version}"
	help       = "sonic {-h --help}"
)

// Help and error help is what is being displaied by the sonic when the an error
// happens for more info refer to the stdlib package flag

var Help string = fmt.Sprintf("%s\n%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n",
	usage, operations, install, remove, query, sync, deptest, version, help)

var ErrHelp = errors.New(Help)
