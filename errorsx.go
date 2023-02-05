// Package errorsx provides extensions for the standard [errors] package.
package errorsx

import (
	"errors"
	"io"
)

// Sentinel is a truly immutable error: unlike errors created via [errors.New],
// it can be declared as a constant.
type Sentinel string

// Error implements the error interface.
func (s Sentinel) Error() string { return string(s) }

// IsAny is a multi-target version of [errors.Is]. See its documentation for
// details.
func IsAny(err, target error, targets ...error) bool {
	for _, t := range append([]error{target}, targets...) {
		if errors.Is(err, t) {
			return true
		}
	}
	return false
}

// AsAny is a multi-target version of [errors.As]. See its documentation for
// details.
func AsAny(err error, target any, targets ...any) bool {
	for _, t := range append([]any{target}, targets...) {
		if errors.As(err, t) {
			return true
		}
	}
	return false
}

// IsTimeout reports whether the error was caused by timeout. Unlike
// [os.IsTimeout], it respects error wrapping.
func IsTimeout(err error) bool {
	var t interface {
		Timeout() bool
	}
	return errors.As(err, &t) && t.Timeout()
}

// Close attempts to close the given [io.Closer] and assigns the returned error (if any) to err.
// If err is already not nil, it will be joined with the [io.Closer]'s error.
//
// NOTE: Close is designed to be used ONLY as a defer statement.
func Close(c io.Closer, err *error) { //nolint:gocritic // ptrToRefParam: false-positive
	if cerr := c.Close(); cerr != nil {
		*err = errors.Join(*err, cerr)
	}
}
