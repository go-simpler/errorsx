// Package errorsx implements extensions for the standard [errors] package.
package errorsx

import "errors"

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

// Do calls the given function and joins the returned error (if any) with err.
// Do is meant to be used in defer statements to record errors returned by [os.File.Close], [sql.Tx.Rollback], etc.
// Note that the returned error in the function where Do is deferred must be named.
func Do(fn func() error, err *error) { //nolint:gocritic // ptrToRefParam: err must be a pointer here.
	*err = errors.Join(*err, fn())
}
