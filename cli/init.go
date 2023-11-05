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
	profile, alias = initialise.Prompt()
	writeOutput(alias, profile)
}

func writeOutput(alias string, profile string) {
	_, err := os.OpenFile(profile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	shellAlias := fmt.Sprintf(`
function %s() {
 if [ "$1" = "set" ]; then
	eval $(%s "$@")
	if [[ "${KEY}" == *"${2}"* ]]; then
		echo -e "\033[0;32m${2} key successfully set!"
	fi
 else
	%[2]s "$@"
 fi
}
`, alias, utils.GetAppDirectory(), shell)
	fmt.Printf(shellAlias)
	// f.WriteString(output)

	f, _ := os.OpenFile(profile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	_, err2 := f.WriteString(shellAlias)
	if err2 != nil {
		log.Fatal(err2)
	}
}
