package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const numFiles = 15
const bytesOfData = 300000
const directory = "/tmp/files"

func outputFileContent(w io.Writer) error {
	// Write out the bytes...
	for i := 0; i < bytesOfData; i++ {
		_, err := w.Write([]byte("0"))
		if err != nil {
			return fmt.Errorf("output error: %v", err)
		}
	}
	return nil
}

func slowOutput(idx int) {
	fName := fmt.Sprintf("%s/%d.txt", directory, idx)
	// Create the file...
	f, err := os.Create(fName)
	// Handle error if cannot create file...
	if err != nil {
		log.Printf("Error creating %q - error: %v", fName, err)
		return
	}
	defer f.Close()
	err = outputFileContent(f)
	if err != nil {
		log.Printf("File content output error: %v", err)
	}
}

func slowOutputToFiles() {
	for i := 0; i < numFiles; i++ {
		slowOutput(i)
	}
}

func main() {
	slowOutputToFiles()
}
