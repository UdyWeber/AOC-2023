package utils

import (
	"log"
	"os"
)

func FileReader(fileName string) string {
	data, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal("Couldn't read file")
	}

	return string(data)
}
