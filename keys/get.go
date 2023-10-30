package keys

import (
	"fmt"
	"keyselector/utils"
	"os"
)

func GetKey(key string) (os.FileInfo, error) {
	path := fmt.Sprintf("%s/%s", utils.GetWorkingDirectory(), key)
	return os.Stat(path)
}
