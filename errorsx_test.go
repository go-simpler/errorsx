package errorsx_test

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"testing"

	"go-simpler.org/errorsx"
)

func TestIsAny(t *testing.T) {
	test := func(name string, err error, targets []error, want bool) {
		t.Helper()
		t.Run(name, func(t *testing.T) {
			t.Helper()
			if got := errorsx.IsAny(err, targets[0], targets[1:]...); got != want {
				t.Errorf("got %t; want %t", got, want)
			}
		})
	}

	test("no matches", errFoo, []error{errBar}, false)
	test("single target match", errFoo, []error{errFoo}, true)
	test("single target match (wrapped)", wrap(errFoo), []error{errFoo}, true)
	test("multiple targets match (wrapped)", wrap(errFoo), []error{errFoo, errBar}, true)
}

func TestHasType(t *testing.T) {
	test := func(name string, fn func(error) bool, err error, want bool) {
		t.Helper()
		t.Run(name, func(t *testing.T) {
			t.Helper()
			if got := fn(err); got != want {
				t.Errorf("got %t; want %t", got, want)
			}
		})
	}

	test("no match", errorsx.HasType[barError], errFoo, false)
	test("match (exact)", errorsx.HasType[fooError], errFoo, true)
	test("match (wrapped)", errorsx.HasType[fooError], wrap(errFoo), true)
	test("match (interface)", errorsx.HasType[interface{ Timeout() bool }], errTimeout, true)
}

func TestSplit(t *testing.T) {
	test := func(name string, err error, wantErrs []error) {
		t.Helper()
		t.Run(name, func(t *testing.T) {
			t.Helper()
			if gotErrs := errorsx.Split(err); !slices.Equal(gotErrs, wantErrs) {
				t.Errorf("got %v; want %v", gotErrs, wantErrs)
			}
		})
	}

	test("nil error", nil, nil)
	test("single error", errFoo, nil)
	test("joined errors (errors.Join)", errors.Join(errFoo, errBar), []error{errFoo, errBar})
	test("joined errors (fmt.Errorf)", fmt.Errorf("%w; %w", errFoo, errBar), []error{errFoo, errBar})
}

func TestIsTimeout(t *testing.T) {
	test := func(name string, fn func(error) bool, err error, want bool) {
		t.Helper()
		t.Run(name, func(t *testing.T) {
			t.Helper()
			if got := fn(err); got != want {
				t.Errorf("got %t; want %t", got, want)
			}
		})
	}

	test("os.IsTimeout", os.IsTimeout, errTimeout, true)
	test("os.IsTimeout (wrapped)", os.IsTimeout, wrap(errTimeout), false)
	test("errorsx.IsTimeout", errorsx.IsTimeout, errTimeout, true)
	test("errorsx.IsTimeout (wrapped)", errorsx.IsTimeout, wrap(errTimeout), true)
}

func TestClose(t *testing.T) {
	test := func(name string, mainErr, closeErr error, wantErrs []error) {
		t.Helper()
		t.Run(name, func(t *testing.T) {
			t.Helper()
			gotErr := func() (err error) {
				c := errCloser{err: closeErr}
				defer errorsx.Close(&c, &err)
				return mainErr
			}()
			for _, wantErr := range wantErrs {
				if !errors.Is(gotErr, wantErr) {
					t.Errorf("got %v; want %v", gotErr, wantErrs)
				}
			}
		})
	}

	test("main: ok; close: ok", nil, nil, []error{})
	test("main: ok; close: error", nil, errBar, []error{errBar})
	test("main: error; close: ok", errFoo, nil, []error{errFoo})
	test("main: error; close: error", errFoo, errBar, []error{errFoo, errBar})
}

var (
	errFoo     fooError
	errBar     barError
	errTimeout timeoutError
)

type fooError struct{}

func (fooError) Error() string { return "foo" }

type barError struct{}

func (barError) Error() string { return "bar" }

type timeoutError struct{}

func (timeoutError) Error() string { return "timeout" }
func (timeoutError) Timeout() bool { return true }

func wrap(err error) error { return fmt.Errorf("wrapped: %w", err) }

type errCloser struct{ err error }

func (c *errCloser) Close() error { return c.err }
