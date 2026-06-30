package main

import (
	"HoldMyBeer/internal/cli"
	"os"
)

// Entry point for the HoldMyBeer application
// Accepts command-line arguments and delegates execution to the CLI package.
func main() {
	os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))
}
