package base63

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_encoder_EncodeDecodeCycle(t *testing.T) {
	count := 1000
	init := 15962909465110768
	e := New()
	for i := init; i <= init+count; i++ {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			enc := e.Encode(uint64(i))
			dec, err := e.Decode(enc)

			assert.NoError(t, err)
			assert.Equal(t, uint64(i), dec)
		})
	}
}

func Test_encoder_Encode(t *testing.T) {
	tests := []struct {
		name string
		num  uint64
		want string
	}{
		{
			name: "zero",
			num:  0,
			want: "0",
		},
		{
			name: "one",
			num:  1,
			want: "1",
		},
		{
			name: "574",
			num:  574,
			want: "97",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := New()

			got := e.Encode(tt.num)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_encoder_Decode(t *testing.T) {
	tests := []struct {
		name               string
		s                  string
		want               uint64
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "zero",
			s:                  "0",
			want:               0,
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "one",
			s:                  "1",
			want:               1,
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "97",
			s:                  "97",
			want:               574,
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "97 with leading zeros",
			s:                  "0097",
			want:               574,
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "empty",
			s:                  "",
			want:               0,
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "not valid symbols",
			s:                  "123!",
			want:               0,
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "very big",
			s:                  "11kyynHVyh",
			want:               15962909465120768,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := New()

			got, err := e.Decode(tt.s)

			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_encoder_validate(t *testing.T) {
	tests := []struct {
		name               string
		s                  string
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "empty",
			s:                  "",
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "not valid symbols",
			s:                  "123!",
			errorAssertionFunc: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := New()

			err := e.validate(tt.s)

			tt.errorAssertionFunc(t, err)
		})
	}
}

func BenchmarkEncoder_Encode(b *testing.B) {
	b.StopTimer()

	id := uint64(15805292566548480)
	e := New()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e.Encode(id)
		id++
	}
}
