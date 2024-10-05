package data

import (
	"math/rand"
	"time"
)

var GlobalRand *rand.Rand

func init() {
	seed := time.Now().UnixNano()
	GlobalRand = rand.New(rand.NewSource(seed))
}
