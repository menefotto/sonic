package limiter

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestLimiter(t *testing.T) {
	f := func() {
		fmt.Println("Random num: ", rand.Intn(99))
	}

	for i := 0; i < 40; i++ {
		Limiter(f, 1000)
	}
}
