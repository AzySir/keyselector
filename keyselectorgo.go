package main

import (
	"fmt"
	"keyselectorgo/cli"
	"log"
	"os"

	"github.com/TwiN/go-color"
)

func main() {
	// log.Println("Running...")
	args := os.Args
	realMain(args)
}

func realMain(args []string) {
	if len(args) <= 1 {
		cli.Help()
		os.Exit(0)
	}
	switch args[1] {
	case "list":
		cli.List()
	case "set":
		if len(args) <= 2 {
			log.SetFlags(0)
			log.Println(color.Red + "\n[ERROR] Please provide a SSH Key name")
			log.Fatalf(
				`
Example:
	keyselector set personal_ssh

				`)
		}
		fmt.Printf("export KEY=%s", cli.Set(args[2]))
	default:
		cli.Help()
	}
}
