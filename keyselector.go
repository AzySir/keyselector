package main

import (
	"keyselector/cli"
	"os"
)

func main() {
	args := os.Args
	realMain(args)
}

func realMain(args []string) {
	if len(args) <= 1 {
		cli.Help()
		os.Exit(0)
	}
	switch args[1] {
	case "autocomplete":
		cli.Autocomplete()
	case "init":
		cli.Init()
	case "list":
		cli.List()
	case "set":
		cli.Set(args)
	case "get":
		cli.Get(args)
	default:
		cli.Help()
	}
}
