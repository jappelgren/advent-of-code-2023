package main

import (
	"log"
	"os"
)

func ParseFile (path string) []byte {
	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatal((err))
	}

	return f
}