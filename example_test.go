package errorsx_test

import (
	"os"

	"github.com/junk1tm/errorsx"
)

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
