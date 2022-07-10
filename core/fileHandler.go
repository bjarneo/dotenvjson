package core

import (
	"log"
	"os"
)

func FileContent(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	return string(data)
}

func WriteFile(filename string, content string) {
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	f.WriteString(content)
}
