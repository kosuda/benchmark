package mysql

import (
	_ "github.com/lib/pq"
	"math/rand"
	"testing"
)

func TestRemove(t *testing.T) {
	Remove()
}

func BenchmarkWrite(b *testing.B) {
	b.ResetTimer()

	for i := 1; i < b.N; i++ {
		point := rand.Intn(100000000000)
		Write(i, point)
	}
}

func BenchmarkRead(b *testing.B) {
	b.ResetTimer()

	for i := 1; i < b.N; i++ {
		Read(i)
	}
}
