package errorsx_test

import (
	"os"

	"github.com/go-simpler/errorsx"
)

//nolint:unused // unused EOF is ok
func ExampleSentinel() {
	const EOF = errorsx.Sentinel("EOF")
}

var err error

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
		defer errorsx.Close(f, &err)

		return nil
	}()
}
