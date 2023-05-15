package internal

import (
	"math/rand"
	"time"
)

func GenerateRandomCode() int {
	min := 1000
	max := 9999
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
