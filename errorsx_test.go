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
		"no match":                         {err: errFoo, targets: []error{errBar}, want: false},
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
		err  error
		as   func(error) bool
		want bool
	}{
		"no match":        {err: errFoo, as: func(err error) bool { return isok(errorsx.As[barError](err)) }, want: false},
		"match (exact)":   {err: errFoo, as: func(err error) bool { return isok(errorsx.As[fooError](err)) }, want: true},
		"match (wrapped)": {err: wrap(errFoo), as: func(err error) bool { return isok(errorsx.As[fooError](err)) }, want: true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := test.as(test.err); got != test.want {
				t.Errorf("got %t; want %t", got, test.want)
			}
		})
	}
}

func TestDo(t *testing.T) {
	tests := map[string]struct {
		mainErr  error
		deferErr error
		wantErrs []error
	}{
		"main: ok; defer: ok":       {mainErr: nil, deferErr: nil, wantErrs: []error{}},
		"main: ok; defer: error":    {mainErr: nil, deferErr: errBar, wantErrs: []error{errBar}},
		"main: error; defer: ok":    {mainErr: errFoo, deferErr: nil, wantErrs: []error{errFoo}},
		"main: error; defer: error": {mainErr: errFoo, deferErr: errBar, wantErrs: []error{errFoo, errBar}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			gotErr := func() (err error) {
				fn := func() error { return test.deferErr }
				defer errorsx.Do(fn, &err)
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

func wrap(err error) error { return fmt.Errorf("%w", err) }
