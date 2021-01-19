package utils

import (
	"math/rand"
	"time"
)

var r *rand.Rand

func SeedRandom() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r = rand.New(s1)
}

func RandInt(n int) int {
	return r.Intn(n)
}
