package main

import (
	"os"
	"testing"
)

func BenchmarkProcessLogFile(b *testing.B) {
	content, err := os.ReadFile("log.txt")
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		lp := NewLogProcessor()
		lp.processLogFile(content)
	}
}
