package util

import "testing"

func BenchmarkInclude(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Includes([]string{"a", "b", "c"}, "a")
	}
}
