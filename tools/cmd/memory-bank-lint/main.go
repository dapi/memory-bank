package main

import (
	"os"

	"github.com/dapi/memory-bank/tools/internal/cli"
)

var version = "dev"

func main() {
	os.Exit(cli.RunLint(os.Args[1:], "memory-bank-lint", version, os.Stdout, os.Stderr))
}
