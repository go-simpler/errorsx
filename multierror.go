package errorsx

import (
	"errors"
	"strings"
)

func Append(err error, errs ...error) error {
	// nothing to append, return err as is.
	if len(errs) == 0 {
		return err
	}

	// remove nil errors if any.
	for i, err := range errs {
		if err == nil {
			errs = append(errs[:i], errs[i+1:]...)
		}
	}

	if err == nil {
		if len(errs) == 1 {
			return errs[0]
		}
		return &multierror{errs: flatten(errs)}
	}

	all := append([]error{err}, errs...)
	return &multierror{errs: flatten(all)}
}

func flatten(errs []error) []error {
	var all []error
	for _, err := range errs {
		// flatten only if the last error in the chain is multierror,
		// otherwise the wrap context might be lost.
		if merr, ok := err.(*multierror); ok {
			all = append(all, merr.errs...)
		} else {
			all = append(all, err)
		}
	}
	return all
}

type multierror struct {
	errs []error
}

// Error implements the error interface.
// TODO(junk1tm): support format customization
func (e *multierror) Error() string {
	msgs := make([]string, len(e.errs))
	for i, err := range e.errs {
		msgs[i] = err.Error()
	}
	return "[" + strings.Join(msgs, " | ") + "]"
}

// Errors provides access to the internal slice of errors.
func (e *multierror) Errors() []error {
	return e.errs
}

// Is implements the non-exported interface used by [errors.Is].
func (e *multierror) Is(target error) bool {
	for _, err := range e.errs {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}

// As implements the non-exported interface used by [errors.As].
func (e *multierror) As(target any) bool {
	for _, err := range e.errs {
		if errors.As(err, target) {
			return true
		}
	}
	return false
}
