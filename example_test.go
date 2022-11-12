package errorsx_test

import (
	"errors"
	"fmt"
	"os"

	"github.com/junk1tm/errorsx"
)

var err error

//nolint:unused // unused EOF is ok
func ExampleSentinel() {
	const EOF = errorsx.Sentinel("EOF")
}

func ExampleIsAny() {
	if errorsx.IsAny(err, os.ErrNotExist, os.ErrPermission) {
		// handle error
	}
}

func ExampleAsAny() {
	if errorsx.AsAny(err, new(*os.PathError), new(*os.LinkError)) {
		// handle error
	}
}

func ExampleIsTimeout() {
	if errorsx.IsTimeout(err) {
		// handle timeout
	}
}

func ExampleClose() {
	_ = func() (err error) {
		f, err := os.Open("file.txt")
		if err != nil {
			return err
		}
		defer errorsx.Close(&err, f) // OR errorsx.Close(&err, f, "closing file: %w")

		return nil
	}()
}

func ExampleAppend() {
	var (
		errFoo = errors.New("foo error")
		errBar = errors.New("bar error")
	)

	print := func(err error) {
		fmt.Printf("%[1]T: %[1]v\n", err)
	}

	print(errorsx.Append(nil))
	print(errorsx.Append(nil, errFoo))
	print(errorsx.Append(nil, errFoo, errBar))

	// Output:
	// <nil>: <nil>
	// *errors.errorString: foo error
	// *errorsx.multierror: [foo error | bar error]
}
