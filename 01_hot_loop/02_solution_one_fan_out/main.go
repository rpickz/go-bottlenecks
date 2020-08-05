package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
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
	var wg sync.WaitGroup
	for i := 0; i < numFiles; i++ {
		i := i // Range variable capture
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
