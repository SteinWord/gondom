package gondom

import (
	"math/rand"
)

const (
	_rawLetters = "abcdefghijklmnopqrstuvwxyz0123456789"
)

func Make(n int, seed int64) string {
	b := make([]byte, n)
	rand.Seed(seed)
	for i := range b {
		b[i] = _rawLetters[rand.Intn(36)]
	}
	return string(b)
}
