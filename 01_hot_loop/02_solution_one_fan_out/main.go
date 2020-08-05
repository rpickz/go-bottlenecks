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

// slowOutput outputs 1000000 bytes of random values to a file whose filename is based on the idx parameter.
func slowOutput(idx int) {
	fName := fmt.Sprintf("%s/%d.txt", directory, idx)
	// Create the file...
	f, err := os.Create(fName)
	// Handle error if cannot create file...
	if err != nil {
		log.Printf("Could not open output file %q for writing - error: %v", fName, err)
		return
	}
	defer f.Close()
	err = outputFileContent(f)
	if err != nil {
		log.Printf("Could not output file contents to file %q - error: %v", fName, err)
	}
}

// slowlyOutputToFiles outputs many files using the slowOutput function.
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
