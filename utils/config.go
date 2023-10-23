package utils

import (
	"fmt"
	"os"
)

var (
	WORKING_DIRECTORY = fmt.Sprintf("%s/.ssh/", os.Getenv("HOME"))
)

func GetWorkingDirectory() string {
	return WORKING_DIRECTORY
}
