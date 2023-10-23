package table

import (
	"io/fs"

	"github.com/jedib0t/go-pretty/table"
)

func PopulateTable(files []fs.DirEntry) table.Writer {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"#", "SSH Key", "Scope", "Description", "Date Created", "Comments"})
	for idx, file := range files {
		t.AppendRow(table.Row{idx + 1, file.Name()})
	}
	t.AppendFooter(table.Row{"Total", len(files) + 1})
	return t
}
