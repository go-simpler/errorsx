// Package errorsx provides extensions for the standard [errors] package.
package errorsx

import (
	"errors"
	"fmt"
	"io"
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

// Close attempts to close the given [io.Closer] and writes the returned error
// (if any) to err. Before being written, the error will be wrapped via
// [fmt.Errorf], the %w verb will be added automatically, there is no need to
// include it in the format.
//
// NOTE: Close is designed to be used ONLY as a defer statement.
func Close(err *error, closer io.Closer, format string, args ...any) {
	if *err != nil {
		// there is already an error, do not override it.
		// TODO(junk1tm): replace with multierror when #1 is closed.
		return
	}
	if cerr := closer.Close(); cerr != nil {
		*err = fmt.Errorf(format+": %w", append(args, cerr)...)
	}
}
