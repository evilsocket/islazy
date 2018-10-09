package async

import (
	"time"
)

// TimedCallback represents a generic function with a return value.
type TimedCallback func() interface{}

// WithTimeout will execute the callback and return its value or a
// TimeoutError if its execution will exceed the provided duration.
func WithTimeout(tm time.Duration, cb TimedCallback) (error, interface{}) {
	timeout := time.After(tm)
	done := make(chan interface{})
	go func() {
		done <- cb()
	}()

	select {
	case <-timeout:
		return TimeoutError, nil
	case res := <-done:
		if res != nil {
			if e, ok := res.(error); ok {
				return e, nil
			}
		}
		return nil, res
	}
}
