package app

import "fmt"

const (
	maxOrigLength  = 2048
	maxShortLength = 16
)

func (sh *Shortener) validateOrig(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("empty string")
	}
	if len(s) > maxOrigLength {
		return fmt.Errorf("string is too long, max len =  %v, get string with len = %v", maxOrigLength, len(s))
	}
	return nil
}

func (sh *Shortener) validateShort(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("empty string")
	}
	if len(s) > maxShortLength {
		return fmt.Errorf("string is too long, max len =  %v, get string with len = %v", maxShortLength, len(s))
	}
	return nil
}
