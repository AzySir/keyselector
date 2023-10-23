package cli

import "log"

func Help() {
	log.SetFlags(0)
	log.Println(`
List of Commands - 

  Options:

  help - List Commands
  list - List of SSH Keys
  metadata - Set MetaData for SSH Key
	`)
}
