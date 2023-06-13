package base63

import (
	"fmt"
)

type Encoder struct {
	power uint64
	alph  string
	codes map[byte]uint64
}

func New() Encoder {
	alph := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	codes := make(map[byte]uint64)
	for i, b := range alph {
		codes[byte(b)] = uint64(i)
	}
	return Encoder{
		power: 63,
		alph:  alph,
		codes: codes,
	}
}

func (e Encoder) Encode(num uint64) string {
	if num == 0 {
		return string(e.alph[0])
	}
	bytes := make([]byte, 0, 10)
	for i := 0; num > 0; i++ {
		ost := num % e.power
		num /= e.power
		bytes = append(bytes, e.alph[ost])
	}
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}

func (e Encoder) Decode(s string) (uint64, error) {
	err := e.validate(s)
	if err != nil {
		return 0, fmt.Errorf("can't validate encoded string: %v", err)
	}
	ret := uint64(0)
	for i := 0; i < len(s); i++ {
		num := e.codes[s[i]]
		ret += num * e.intPow(e.power, len(s)-1-i)
	}
	return ret, nil
}

func (e Encoder) intPow(n uint64, m int) uint64 {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func (e Encoder) validate(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("empty encoded string")
	}
	for _, b := range s {
		_, ok := e.codes[byte(b)]
		if !ok {
			return fmt.Errorf("no such symbol in alph: %v", b)
		}
	}
	return nil
}
