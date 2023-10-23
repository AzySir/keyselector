package utils

import (
	"io/fs"
	"log"
	"os"
)

func ReadKeys() []fs.DirEntry {
	log.Println(GetWorkingDirectory())
	files, err := os.ReadDir(GetWorkingDirectory())
	if err != nil {
		log.Printf("[ERROR] Fetching Keys - ReadKeys()")
		log.Println(err)
		log.Fatalf("[ERROR] %s", err)
	}
	return files
}
