//Package limiter implements a single function which does implement the decorator
// pattern , what is does is limits the number of go routines spanned, the number
// of concurrent go routines is controlles by the exported variable named GorountineNum
// it does not take care of error handling in any way. The limiter function accept
// a function returning a value and and errror named of type LimiterFunc, and a sleep
// time wich is inserted right after a function call.

package limiter

import (
	"runtime"
	"time"
)

var (
	GoroutineNum = runtime.GOMAXPROCS(0) * 2
	limiterChan  = make(chan struct{}, GoroutineNum)
)

func Limiter(fn func(), wait time.Duration) {
	limiterChan <- struct{}{}
	go func() {
		defer func() {
			<-limiterChan
		}()
		fn()

		time.Sleep(wait)
	}()
}
