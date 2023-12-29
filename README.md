# errorsx

[![checks](https://github.com/go-simpler/errorsx/actions/workflows/checks.yml/badge.svg)](https://github.com/go-simpler/errorsx/actions/workflows/checks.yml)
[![pkg.go.dev](https://pkg.go.dev/badge/go-simpler.org/errorsx.svg)](https://pkg.go.dev/go-simpler.org/errorsx)
[![goreportcard](https://goreportcard.com/badge/go-simpler.org/errorsx)](https://goreportcard.com/report/go-simpler.org/errorsx)
[![codecov](https://codecov.io/gh/go-simpler/errorsx/branch/main/graph/badge.svg)](https://codecov.io/gh/go-simpler/errorsx)

Extensions for the standard `errors` package.

## 📦 Install

Go 1.21+

```shell
go get go-simpler.org/errorsx
```

## 📋 Usage

### IsAny

A multi-target version of `errors.Is`.

```go
if errorsx.IsAny(err, os.ErrNotExist, os.ErrPermission) {
    fmt.Println(err)
}
```

### HasType

Reports whether the error has type `T`.
It is equivalent to `errors.As` without the need to declare the target variable.

```go
if errorsx.HasType[*os.PathError](err) {
    fmt.Println(err)
}
```

### Split

Returns errors joined by `errors.Join` or by `fmt.Errorf` with multiple `%w` verbs.
If the given error was created differently, `Split` returns nil.

```go
if errs := errorsx.Split(err); errs != nil {
    fmt.Println(errs)
}
```

### Close

Attempts to close the given `io.Closer` and assigns the returned error (if any) to `err`.

```go
func() (err error) {
    f, err := os.Open("file.txt")
    if err != nil {
        return err
    }
    defer errorsx.Close(f, &err)

    return nil
}()
```
