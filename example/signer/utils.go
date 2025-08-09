package main

import (
	"io"
	"log"
	"os"
)

func writeFile(buf []byte, path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.Write(buf)
	if err != nil {
		log.Fatal(err)
	}
}

func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return buf, nil
}
