package utils

import (
	"fmt"
	"log"
	"os"
)

var (
	WORKING_DIRECTORY = fmt.Sprintf("%s/.ssh/", os.Getenv("HOME"))
)

func GetWorkingDirectory() string {
	return WORKING_DIRECTORY
}

func GetAppDirectory() string {
	path, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	return path
}
