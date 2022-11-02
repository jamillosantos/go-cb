package cb

import (
	"errors"

	"github.com/sony/gobreaker"
)

var (
	ErrInvalidConversion = errors.New("invalid conversion")
)

// Execute executes the given function and returns the result using generics so no conversion need to be handled
// by the caller.
func Execute[T any](cb *gobreaker.CircuitBreaker, f func() (T, error)) (r T, rerr error) {
	anyR, err := cb.Execute(func() (interface{}, error) {
		r, err := f()
		return r, err
	})
	if err != nil {
		rerr = err
		return
	}
	return anyR.(T), nil
}
