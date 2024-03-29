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

// HasType reports whether the error has type T.
// It is equivalent to [errors.As] without the need to declare the target variable.
func HasType[T any](err error) bool {
	var t T
	return errors.As(err, &t)
}

// Split returns errors joined by [errors.Join] or by [fmt.Errorf] with multiple %w verbs.
// If the given error was created differently, Split returns nil.
func Split(err error) []error {
	u, ok := err.(interface{ Unwrap() []error })
	if !ok {
		return nil
	}
	return u.Unwrap()
}

// Close attempts to close the given [io.Closer] and assigns the returned error (if any) to err.
// If err is already not nil, it will be joined with the [io.Closer]'s error.
func Close(c io.Closer, err *error) { //nolint:gocritic // ptrToRefParam: err must be a pointer here.
	if cerr := c.Close(); cerr != nil {
		*err = errors.Join(*err, cerr)
	}
}
