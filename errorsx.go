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

// Close attempts to close the given [io.Closer] and assigns the returned error
// (if any) to err. If optional formatAndArgs are provided, the error will be
// wrapped via [fmt.Errorf] before being assigned. Do not include err in
// formatAndArgs, it will be appended automatically.
//
// NOTE: Close is designed to be used ONLY as a defer statement.
func Close(err *error, closer io.Closer, formatAndArgs ...any) {
	if *err != nil {
		// there is already an error, do not override it.
		// TODO(junk1tm): replace with multierror when #1 is closed.
		return
	}
	if cerr := closer.Close(); cerr != nil {
		if len(formatAndArgs) > 0 {
			format, args := formatAndArgs[0].(string), formatAndArgs[1:]
			*err = fmt.Errorf(format, append(args, cerr)...)
		} else {
			*err = cerr
		}
	}
}
