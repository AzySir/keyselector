package cli

import (
	"keyselectorgo/keys"
	"log"
)

func Set(key string) string {
	_, err := keys.GetKey(key)
	if err != nil {
		log.Fatalf("[ERROR] %s does NOT exist", key)
	}
	return key
}
