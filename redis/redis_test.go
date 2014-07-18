package redis

import (
	"testing"
)

const (
	max = 100000
)

func TestClear(t *testing.T) {
	Clear()

	if !Empty() {
		t.Error("Not be cleared")
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
