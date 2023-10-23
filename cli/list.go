package cli

import (
	"keyselectorgo/table"
	"keyselectorgo/utils"
	"log"
)

func List() {
	log.SetFlags(0)
	keys := utils.ReadKeys()
	t := table.PopulateTable(keys)
	log.Println(t.Render())
}
