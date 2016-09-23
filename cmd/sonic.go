package main

import (
	"fmt"
	"os"

	"github.com/sonic/cmd/cmdflags"
	"github.com/sonic/lib/confparse"
	"github.com/sonic/lib/operations"
)

func main() {
	cmdflags.Operations.InstallPkgs = operations.Install
	cmdflags.Operations.RemovePkgs = operations.Remove
	cmdflags.Operations.DeptestPkgs = operations.DepTest
	cmdflags.Operations.QueryPkgs = operations.Query
	cmdflags.Operations.SyncDbs = operations.Sync

	parser, err := confparse.NewParserFromFile("lib/operations/sonic.conf")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	err = cmdflags.CommandExecutor(parser)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in sonic.go", r)
		}
	}()
}
