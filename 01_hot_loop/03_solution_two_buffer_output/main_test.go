package main

import "testing"

var err error
var content []byte

func BenchmarkGenerateFileContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		content, err = outputFileContent()
	}
}
