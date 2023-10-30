package cli

import (
	"bufio"
	"log"
	"os"
)

func Autocomplete() byte {
	consoleReader := bufio.NewReaderSize(os.Stdin, 1)
	key, err := consoleReader.ReadByte()
	if err != nil {
		log.Fatalf("Error")
	}
	log.Println(key)
	return key
}
