package id_gen

import (
	"testing"
)

func BenchmarkIDGenerator_Generate(b *testing.B) {
	b.StopTimer()

	generator, err := New(1)
	if err != nil {
		b.Fatalf("can't create id generator")
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = generator.Generate()
	}
}
