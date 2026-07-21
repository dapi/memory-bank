package main

import (
	"os"

	"github.com/dapi/memory-bank/tools/internal/cli"
)

var version = "dev"

func main() {
	os.Exit(cli.Run(os.Args[1:], version, os.Stdout, os.Stderr))
}
