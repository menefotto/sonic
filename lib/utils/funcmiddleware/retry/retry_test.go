package retry

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
)

func TestRetry(t *testing.T) {
	var fn RetryFunc = func() (interface{}, error) {
		res := rand.Intn(5)
		if res == 4 {
			return "ok", nil
		}
		fmt.Println("Failed retrying now")

		return "", errors.New("test error")
	}

	result, err := Retry(fn, 5, 1000)
	if err != nil {
		t.Fatal("impossible")
	}

	t.Logf("Result is as expected is: %s\n", result)
}
