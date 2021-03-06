package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sync"
)

const numFiles = 15
const bytesOfData = 300000
const directory = "/tmp/files"

func outputFileContent() ([]byte, error) {
	var buf bytes.Buffer
	// Write out the bytes...
	for i := 0; i < bytesOfData; i++ {
		_, err := buf.WriteRune('0')
		if err != nil {
			return nil, fmt.Errorf("output error: %v", err)
		}
	}
	return buf.Bytes(), nil
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

	content, err := outputFileContent()
	if err != nil {
		log.Printf("Content generation error: %v", err)
		return
	}

	_, err = f.Write(content)
	if err != nil {
		log.Printf("File output error: %v", err)
	}
}

func slowOutputToFiles() {
	var wg sync.WaitGroup
	for i := 0; i < numFiles; i++ {
		i := i
		wg.Add(1)
		go func() {
			slowOutput(i)
			wg.Done()
		}()
	}

	wg.Wait()
}

func main() {
	slowOutputToFiles()
}
