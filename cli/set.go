package cli

import (
	"fmt"
	"keyselector/keys"
	"keyselector/utils"
	"log"

	"github.com/TwiN/go-color"
)

func Set(args []string) {
	if len(args) <= 2 {
		log.SetFlags(0)
		log.Println(color.Red + "\n[ERROR] Please provide a SSH Key name")
		log.Fatalf(
			`
Example:
keyselector set personal_ssh

			`)
	}

	output(args[2])
}

func output(key string) {
	_, err := keys.GetKey(key)
	if err != nil {
		log.Fatalf("[ERROR] %s does NOT exist", key)
	}
	fmt.Printf("export KEY=%s%s", utils.GetWorkingDirectory(), key)
}
