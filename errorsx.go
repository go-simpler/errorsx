// Package errorsx provides extensions for the standard [errors] package.
package errorsx

import (
	"errors"
)

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
