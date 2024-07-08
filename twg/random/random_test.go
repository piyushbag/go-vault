package random

import (
	"math/rand"
	"testing"
	"time"
)

func TestPick(t *testing.T) {
	seed := time.Unix(0, 0).UnixNano()
	r := rand.New(rand.NewSource(seed))
	arg := make([]int, 10)
	for i := 0; i < 10; i++ {
		arg[i] = r.Int()
	}
	got := Pick(arg)
	for _, v := range arg {
		if got == v {
			return
		}
	}
	t.Errorf("Pick(%v) = %v", seed, got)
}
