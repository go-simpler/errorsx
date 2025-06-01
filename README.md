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
// Before:
if errors.Is(err, os.ErrNotExist) || errors.Is(err, os.ErrPermission) {
    fmt.Println(err)
}

// After:
if errorsx.IsAny(err, os.ErrNotExist, os.ErrPermission) {
    fmt.Println(err)
}
```

### As

A generic version of `errors.As`.

```go
// Before:
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    fmt.Println(pathErr.Path)
}

// After:
if pathErr, ok := errorsx.As[*os.PathError](err); ok {
    fmt.Println(pathErr.Path)
}
```

### Do

Calls the given function and joins the returned error (if any) with `err`.

```go
f, err := os.Open("file.txt")
if err != nil {
    return err
}

// Before:
defer func() {
    err = errors.Join(err, f.Close())
}()

// After:
defer errorsx.Do(f.Close, &err)
```
