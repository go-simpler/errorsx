// Package errorsx provides extensions for the standard [errors] package.
package errorsx

import (
	"errors"
)

// IsOneOf is a multi-target version of [errors.Is]. See its documentation for
// details.
func IsOneOf(err error, targets ...error) bool {
	for _, t := range targets {
		if errors.Is(err, t) {
			return true
		}
	}
	return false
}

// AsOneOf is a multi-target version of [errors.As]. See its documentation for
// details.
func AsOneOf(err error, targets ...any) bool {
	for _, t := range targets {
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
