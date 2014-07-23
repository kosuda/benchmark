package postgres

import (
	_ "github.com/lib/pq"
	"math/rand"
	"testing"
)

var maxID int

func TestRemove(t *testing.T) {
	Remove()
}

func BenchmarkWrite(b *testing.B) {
	b.ResetTimer()

	var id int
	for id = 1; id < b.N; id++ {
		point := rand.Intn(100000000000)
		Write(id, point)
	}

	maxID = id
}

func BenchmarkRead(b *testing.B) {
	b.ResetTimer()

	for i := 1; i < b.N; i++ {
		id := rand.Intn(maxID)
		Read(id)
	}
}
