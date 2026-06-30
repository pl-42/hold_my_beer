package main

import (
	"os"

	cli "github.com/pl-42/hold_my_beer/internal/cli"
)

// Entry point for the HoldMyBeer application
// Accepts command-line arguments and delegates execution to the CLI package.
func main() {
	os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))
}
