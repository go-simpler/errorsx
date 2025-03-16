// Package errorsx implements extensions for the standard [errors] package.
package errorsx

import (
	"errors"
	"io"
)

// IsAny is a multi-target version of [errors.Is].
func IsAny(err, target error, targets ...error) bool {
	if errors.Is(err, target) {
		return true
	}
	for _, t := range targets {
		if errors.Is(err, t) {
			return true
		}
	}
	return false
}

// As is a generic version of [errors.As].
func As[T any](err error) (T, bool) {
	var t T
	ok := errors.As(err, &t)
	return t, ok
}

// Close attempts to close the given [io.Closer] and assigns the returned error (if any) to err.
// If err is already not nil, it will be joined with the [io.Closer]'s error.
func Close(c io.Closer, err *error) { //nolint:gocritic // ptrToRefParam: err must be a pointer here.
	*err = errors.Join(*err, c.Close())
}
