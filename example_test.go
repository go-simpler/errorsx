package errorsx_test

import (
	"fmt"
	"os"

	"go-simpler.org/errorsx"
)

var err error

func ExampleIsAny() {
	if errorsx.IsAny(err, os.ErrNotExist, os.ErrPermission) {
		fmt.Println(err)
	}
}

func ExampleHasType() {
	if errorsx.HasType[*os.PathError](err) {
		fmt.Println(err)
	}
}

func ExampleSplit() {
	if errs := errorsx.Split(err); errs != nil {
		fmt.Println(errs)
	}
}

//nolint:errcheck // this is just an example.
func ExampleClose() {
	func() (err error) {
		f, err := os.Open("file.txt")
		if err != nil {
			return err
		}
		defer errorsx.Close(f, &err)

		return nil
	}()
}
