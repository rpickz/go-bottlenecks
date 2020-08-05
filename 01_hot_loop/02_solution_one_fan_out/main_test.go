package main

import "testing"

var err error

type nopWriter struct{}

func (n nopWriter) Write(content []byte) (int, error) {
	return 0, nil
}

func BenchmarkGenerateFileContent(b *testing.B) {
	w := nopWriter{}
	for i := 0; i < b.N; i++ {
		err = outputFileContent(w)
	}
}
