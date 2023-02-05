# errorsx

[![ci](https://github.com/go-simpler/errorsx/actions/workflows/go.yml/badge.svg)](https://github.com/go-simpler/errorsx/actions/workflows/go.yml)
[![docs](https://pkg.go.dev/badge/github.com/go-simpler/errorsx.svg)](https://pkg.go.dev/github.com/go-simpler/errorsx)
[![report](https://goreportcard.com/badge/github.com/go-simpler/errorsx)](https://goreportcard.com/report/github.com/go-simpler/errorsx)
[![codecov](https://codecov.io/gh/go-simpler/errorsx/branch/main/graph/badge.svg)](https://codecov.io/gh/go-simpler/errorsx)

Extensions for the standard `errors` package

## 📦 Install

```shell
go get github.com/go-simpler/errorsx
```

## 🧩 Extensions

### Sentinel

A truly immutable error: unlike errors created via `errors.New`, it can be declared as a constant.

```go
const EOF = errorsx.Sentinel("EOF")
```

### IsAny

A multi-target version of `errors.Is`.

```go
if errorsx.IsAny(err, os.ErrNotExist, os.ErrPermission) {
	// handle error
}
```

### AsAny

A multi-target version of `errors.As`.

```go
if errorsx.AsAny(err, new(*os.PathError), new(*os.LinkError)) {
	// handle error
}
```

### IsTimeout

Reports whether the error was caused by timeout.
Unlike `os.IsTimeout`, it respects error wrapping.

```go
if errorsx.IsTimeout(err) {
	// handle timeout
}
```

### Close

Attempts to close the given `io.Closer` and assigns the returned error (if any) to `err`.

```go
f, err := os.Open("file.txt")
if err != nil {
	return err
}
defer errorsx.Close(f, &err)
```
