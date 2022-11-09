package errorsx_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/junk1tm/errorsx"
)

type fooError struct{}

func (fooError) Error() string { return "foo" }

type barError struct{}

func (barError) Error() string { return "bar" }

type timeoutError struct{}

func (timeoutError) Error() string { return "timeout" }
func (timeoutError) Timeout() bool { return true }

var (
	errFoo     fooError
	errBar     barError
	errTimeout timeoutError
)

func wrap(err error) error { return fmt.Errorf("wrapped: %w", err) }

func TestIsAny(t *testing.T) {
	test := func(name string, err error, targets []error, want bool) {
		t.Helper()
		t.Run(name, func(t *testing.T) {
			t.Helper()
			if got := errorsx.IsAny(err, targets[0], targets[1:]...); got != want {
				t.Errorf("got %v; want %v", got, want)
			}
		})
	}

	test("no matches", errFoo, []error{errBar}, false)
	test("single target match", errFoo, []error{errFoo}, true)
	test("single target match (wrapped)", wrap(errFoo), []error{errFoo}, true)
	test("multiple targets match (wrapped)", wrap(errFoo), []error{errFoo, errBar}, true)
}

func TestAsAny(t *testing.T) {
	test := func(name string, err error, targets []any, want bool) {
		t.Helper()
		t.Run(name, func(t *testing.T) {
			t.Helper()
			if got := errorsx.AsAny(err, targets[0], targets[1:]...); got != want {
				t.Errorf("got %v; want %v", got, want)
			}
		})
	}

	test("no matches", errFoo, []any{new(barError)}, false)
	test("single target match", errFoo, []any{new(fooError)}, true)
	test("single target match (wrapped)", wrap(errFoo), []any{new(fooError)}, true)
	test("multiple targets match (wrapped)", wrap(errFoo), []any{new(fooError), new(barError)}, true)
}

func TestIsTimeout(t *testing.T) {
	test := func(name string, fn func(error) bool, err error, want bool) {
		t.Helper()
		t.Run(name, func(t *testing.T) {
			t.Helper()
			if got := fn(err); got != want {
				t.Errorf("got %v; want %v", got, want)
			}
		})
	}

	test("os.IsTimeout", os.IsTimeout, errTimeout, true)
	test("os.IsTimeout (wrapped)", os.IsTimeout, wrap(errTimeout), false)
	test("errorsx.IsTimeout", errorsx.IsTimeout, errTimeout, true)
	test("errorsx.IsTimeout (wrapped)", errorsx.IsTimeout, wrap(errTimeout), true)
}

type badFile struct{}

func (*badFile) Close() error { return errors.New("won't close 😈") }

func openFile() (*badFile, error) { return new(badFile), nil }

func testFile() (err error) {
	f, err := openFile()
	if err != nil {
		return fmt.Errorf("opening bad file: %w", err)
	}
	defer errorsx.Close(&err, f, "closing bad file")

	// do something with f

	return nil
}

func TestClose(t *testing.T) {
	if err := testFile(); err.Error() != "closing bad file: won't close 😈" {
		t.Error("unexpected error message")
	}
}
