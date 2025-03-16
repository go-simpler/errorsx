package errorsx_test

import (
	"errors"
	"fmt"
	"testing"

	"go-simpler.org/errorsx"
)

func TestIsAny(t *testing.T) {
	tests := map[string]struct {
		err     error
		targets []error
		want    bool
	}{
		"no matches":                       {err: errFoo, targets: []error{errBar}, want: false},
		"single target match":              {err: errFoo, targets: []error{errFoo}, want: true},
		"single target match (wrapped)":    {err: wrap(errFoo), targets: []error{errFoo}, want: true},
		"multiple targets match (wrapped)": {err: wrap(errFoo), targets: []error{errBar, errFoo}, want: true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := errorsx.IsAny(test.err, test.targets[0], test.targets[1:]...); got != test.want {
				t.Errorf("got %t; want %t", got, test.want)
			}
		})
	}
}

func TestAs(t *testing.T) {
	isok := func(_ any, ok bool) bool { return ok }

	tests := map[string]struct {
		fn   func(error) bool
		err  error
		want bool
	}{
		"no match":        {fn: func(err error) bool { return isok(errorsx.As[barError](err)) }, err: errFoo, want: false},
		"match (exact)":   {fn: func(err error) bool { return isok(errorsx.As[fooError](err)) }, err: errFoo, want: true},
		"match (wrapped)": {fn: func(err error) bool { return isok(errorsx.As[fooError](err)) }, err: wrap(errFoo), want: true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := test.fn(test.err); got != test.want {
				t.Errorf("got %t; want %t", got, test.want)
			}
		})
	}
}

func TestClose(t *testing.T) {
	tests := map[string]struct {
		mainErr  error
		closeErr error
		wantErrs []error
	}{
		"main: ok; close: ok":       {mainErr: nil, closeErr: nil, wantErrs: []error{}},
		"main: ok; close: error":    {mainErr: nil, closeErr: errBar, wantErrs: []error{errBar}},
		"main: error; close: ok":    {mainErr: errFoo, closeErr: nil, wantErrs: []error{errFoo}},
		"main: error; close: error": {mainErr: errFoo, closeErr: errBar, wantErrs: []error{errFoo, errBar}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			gotErr := func() (err error) {
				c := errCloser{err: test.closeErr}
				defer errorsx.Close(&c, &err)
				return test.mainErr
			}()
			for _, wantErr := range test.wantErrs {
				if !errors.Is(gotErr, wantErr) {
					t.Errorf("got %v; want %v", gotErr, wantErr)
				}
			}
		})
	}
}

var (
	errFoo fooError
	errBar barError
)

type fooError struct{}

func (fooError) Error() string { return "foo" }

type barError struct{}

func (barError) Error() string { return "bar" }

type errCloser struct{ err error }

func (c *errCloser) Close() error { return c.err }

func wrap(err error) error { return fmt.Errorf("%w", err) }
