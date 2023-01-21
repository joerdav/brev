package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/joerdav/brev/repl"
)

func run() error {
	version := "devel"
	in, ok := debug.ReadBuildInfo()
	if ok && in.Main.Version != "" {
		version = in.Main.Version
	}
	fmt.Printf("Brev Repl (%s)\n", version)
	repl.Start(os.Stdin, os.Stdout)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
