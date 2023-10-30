package cli

import "log"

func Help() {
	log.SetFlags(0)
	log.Printf(`
List of Commands:

  COMMANDS        PARAMS          DESCRIPTION
  config                          Sets Configuration for Keyselector (alias, shell profile etc.)
  set             <key name>      Sets the key you want to use as the ENV variable KEY
  help                            List Commands
  list                            List of SSH Keys
  metadata                        Set MetaData for SSH Key
	`)
}
