package utils

import (
	"log"
	"os"
)

func OpenFile(fn string) string {
	content, err := os.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
