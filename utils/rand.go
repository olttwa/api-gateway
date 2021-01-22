package utils

import (
	"math/rand"
	"time"
)

var (
	seeded bool
	r      *rand.Rand
)

func RandInt(n int) int {
	if !seeded {
		seeded = true
		s1 := rand.NewSource(time.Now().UnixNano())
		r = rand.New(s1)
	}

	return r.Intn(n)
}
