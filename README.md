# errorsx

[![checks](https://github.com/go-simpler/errorsx/actions/workflows/checks.yml/badge.svg)](https://github.com/go-simpler/errorsx/actions/workflows/checks.yml)
[![pkg.go.dev](https://pkg.go.dev/badge/go-simpler.org/errorsx.svg)](https://pkg.go.dev/go-simpler.org/errorsx)
[![goreportcard](https://goreportcard.com/badge/go-simpler.org/errorsx)](https://goreportcard.com/report/go-simpler.org/errorsx)
[![codecov](https://codecov.io/gh/go-simpler/errorsx/branch/main/graph/badge.svg)](https://codecov.io/gh/go-simpler/errorsx)

Extensions for the standard `errors` package

## 📦 Install

```shell
go get go-simpler.org/errorsx
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

### HasType

Reports whether the error has type `T`.
It is equivalent to `errors.As` without the need to declare the target variable.

```go
if errorsx.HasType[*os.PathError](err) {
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
