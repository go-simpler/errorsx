package errorsx_test

import (
	"os"

	"go-simpler.org/errorsx"
)

var err error

func ExampleIsAny() {
	if errorsx.IsAny(err, os.ErrNotExist, os.ErrPermission) {
		// handle error
	}
}

func ExampleHasType() {
	if errorsx.HasType[*os.PathError](err) {
		// handle error
	}
}

func ExampleSplit() {
	if errs := errorsx.Split(err); errs != nil {
		// handle errors
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
