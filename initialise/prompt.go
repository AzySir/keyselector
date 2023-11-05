package initialise

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const colorWhite = "\033[1;37m"
const colorNone = "\033[0m"
const colorRed = "\033[0;31m"
const colorGreen = "\033[0;32m"
const colorYellow = "\033[1;33m"

func Prompt() (string, string) {
	fmt.Println("Initializing...")
	alias := readAlias()
	profile := lookUpProfile(os.Getenv("SHELL"))

	return profile, alias
}

func readAlias() string {
	fmt.Println("What would you like the alias command to be? (Default: keyselector)")
	alias := "keyselector"
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	if (input != alias) && (input != "") {
		alias = input
	}
	fmt.Printf("%s%s has been selected as alias for usage%s\n", colorGreen, alias, colorNone)
	return alias
}

func lookUpProfile(shell string) string {
	home := os.Getenv("HOME")
	switch shell {
	case "/bin/zsh":
		shell = fmt.Sprintf("%s/.zshrc", home)
	case "/bin/bash":
		shell = fmt.Sprintf("%s/.bashrc", home)
	}
	return shell
}
