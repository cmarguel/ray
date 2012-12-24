package util

import (
	"math/rand"
)

func ShuffleFloat64(deck []float64) {
	perm := rand.Perm(len(deck))
	for i, v := range perm {
		deck[i], deck[v] = deck[v], deck[i]
	}
}
