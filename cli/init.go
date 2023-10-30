package cli

import (
	"fmt"
	"keyselector/initialise"
	"keyselector/utils"
	"log"
	"os"
)

var (
	alias   string
	profile string
	shell   string = "keyselector_shell"
)

func Init() {
	// 1) Prompt for Alias creation
	// 2) Explanation of Alias
	// 3) Creating alias
	// 4) Adding to file
	profile, alias = initialise.Prompt()
	writeOutput(alias, profile)
}

func writeOutput(alias string, profile string) {
	f, err := os.OpenFile(profile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	output := fmt.Sprintf(`
%s() {
 . %s/%s "${@}"
}
`, alias, utils.GetAppDirectory(), shell)

	f.WriteString(output)

}
