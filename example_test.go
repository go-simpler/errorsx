package errorsx_test

import (
	"os"

	"github.com/junk1tm/errorsx"
)

var err error

func ExampleIsOneOf() {
	if errorsx.IsOneOf(err, os.ErrNotExist, os.ErrPermission) {
		// handle error
	}
}

func ExampleAsOneOf() {
	if errorsx.AsOneOf(err, new(*os.PathError), new(*os.LinkError)) {
		// handle error
	}
}

func ExampleIsTimeout() {
	if errorsx.IsTimeout(err) {
		// handle timeout
	}
}
