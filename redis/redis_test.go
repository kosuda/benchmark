package redis

import (
	"testing"
)

const (
	max = 100000
)

func BenchmarkClear(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Clear()
	}
}

func BenchmarkWrite(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Write(max)
	}
}

func BenchmarkRead(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Read(max)
	}
}
