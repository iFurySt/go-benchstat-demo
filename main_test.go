package main

import (
	"os"
	"testing"
)

func BenchmarkBiz(b *testing.B) {
	old := os.Stdout
	devNull, _ := os.Open(os.DevNull)
	defer devNull.Close()
	os.Stdout = devNull
	defer func() { os.Stdout = old }()

	for i := 0; i < b.N; i++ {
		mockBiz()
	}
}
