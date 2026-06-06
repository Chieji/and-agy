package main

import (
	"fmt"
	"os"

	"github.com/Chieji/and-agy/internal/cli"
)

var (
	Version = "dev"
	Commit  = "unknown"
	Date    = "unknown"
)

func main() {
	if len(os.Args) > 0 && os.Args[0] == "version" {
		fmt.Printf("and-agy version %s\n", Version)
		fmt.Printf("commit: %s\n", Commit)
		fmt.Printf("built: %s\n", Date)
		return
	}

	if err := cli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}