package initialise

import (
	"bufio"
	"fmt"
	"log"
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
	profile := readProfile()

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

func promptNewFile(scanner *bufio.Scanner, input string) string {
	fmt.Printf("%s%s does not exist.\n%sWould you like to create this file?\n%s", colorYellow, input, colorWhite, colorNone)
	scanner.Scan()
	selection := scanner.Text()
	if selection == "yes" || selection == "y" || selection == "" {
		f, err := os.Create(input)
		if err != nil {
			defer f.Close()
			log.Fatalf("%s[ERROR] %s", colorRed, err)
		}
		defer f.Close()
	}

	_, err := os.Stat(input)
	if err != nil {
		log.Fatal(err)
	}

	return input
}

func readProfile() string {
	shell := os.Getenv("SHELL")
	profile := lookUpProfile(shell)
	fmt.Printf("\n%sEnter what shell profile you would like to use? (Default: %s)\n\n%sHit enter to continue with default.\n\n", colorWhite, profile, colorNone)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	if input == "" {
		input = profile
		input = strings.Replace(input, "~", os.Getenv("HOME"), -1)
	}
	_, err := os.Stat(input)

	// if File Does not exist - prompt creation
	if err != nil {
		input = strings.Replace(input, "~", os.Getenv("HOME"), -1)
		log.Printf("%s %s", colorYellow, input)
		input = promptNewFile(scanner, input)
		log.Printf("SECOND: %s %s", colorYellow, input)
	}
	return input
}
