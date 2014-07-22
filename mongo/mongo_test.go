package mongo

import (
	"math/rand"
	"testing"
)

var maxID int

func TestRemove(t *testing.T) {
	if !IsEmpty() {
		Remove()
	}

	if !IsEmpty() {
		t.Error("not be cleard")
	}
}

func BenchmarkWrite(b *testing.B) {
	b.ResetTimer()

	id := 0
	for ; id < b.N; id++ {
		point := rand.Intn(100000000000)
		Write(id, point)
	}

	maxID = id
}

func BenchmarkRead(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		id := rand.Intn(maxID)
		Read(id)
	}
}
