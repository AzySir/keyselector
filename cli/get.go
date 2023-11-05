package cli

// Desired Output

// ----------------------------------------------
// | KeyName |	LastUsed | For Data in MetaData: |
// ----------------------------------------------
// Open our jsonFile
import (
	"fmt"
	"keyselector/utils"
	"log"
	"os"
)

type error interface {
	Error() string
}

var accept_files = [...]string{"ks_metadata.json", "metadata.json", "keyselector_metadata.json"}

func Get(args []string) {
	if len(args) < 3 {
		log.Fatal("Please select a key to get information from i.e keyselector get mykey")
	}
	metadataFile, err := getMetadataFile()
	if err != nil {
		log.Fatal(err)
	}
	jsonFile, err := os.Open(metadataFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(metadataFile)

	defer jsonFile.Close()
	// log.Println(json.Unmarshal(byteValue))
}

func getMetadataFile() (string, error) {
	for _, file := range accept_files {
		if utils.FileExists(fmt.Sprintf("%s/.ssh/%s", os.Getenv("HOME"), file)) {
			return file, nil
		}
	}
	return "", fmt.Errorf("metadata files not found")
}
