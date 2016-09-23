// Package retry provides only one function, which implements the decorator patter
// basically wrapping another function and implemeting a retry and wait meccanism
// in case of failure, the number of retries is given by function parameter as well
// as the time to wait before the next retry

package retry

import "time"

type RetryFunc func() (interface{}, error)

func Retry(fn RetryFunc, retries uint, wait time.Duration) (interface{}, error) {

	var (
		result interface{}
		err    error
	)

	for i := retries; i >= 0; i-- {
		result, err := fn()

		if err != nil {
			retries--
			time.Sleep(wait)
		}

		if err == nil {
			return result, nil
		}
	}

	return result, err
}
