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

func ExampleAs() {
	if pathErr, ok := errorsx.As[*os.PathError](err); ok {
		fmt.Println(pathErr.Path)
	}
}

func ExampleDo() {
	_ = func() (err error) {
		f, err := os.Open("file.txt")
		if err != nil {
			return err
		}
		defer errorsx.Do(f.Close, &err)

		return nil
	}
}
