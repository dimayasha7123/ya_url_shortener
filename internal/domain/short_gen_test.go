package domain

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"ya_url_shortener/internal/domain/base63"
	"ya_url_shortener/internal/domain/id_gen"
)

// tests uniq and length
func Test_ShortURLGenerator(t *testing.T) {
	count := 10000
	generator, err := id_gen.New(1)
	require.NoError(t, err)

	encoder := base63.New()

	shorts := make(map[string]struct{}, count)
	for i := 0; i < count; i++ {
		id := generator.Generate()
		enc := encoder.Encode(id)

		assert.Len(t, enc, 10)

		_, ok := shorts[enc]

		assert.False(t, ok)

		shorts[enc] = struct{}{}
	}
}

func Benchmark_ShortURLGenerator(b *testing.B) {
	b.StopTimer()

	generator, err := id_gen.New(1)
	if err != nil {
		b.Fatalf("can't create id generator")
	}
	encoder := base63.New()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = encoder.Encode(generator.Generate())
	}
}
