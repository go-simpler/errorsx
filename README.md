# errorsx

[![ci](https://github.com/junk1tm/errorsx/actions/workflows/go.yml/badge.svg)](https://github.com/junk1tm/errorsx/actions/workflows/go.yml)
[![docs](https://pkg.go.dev/badge/github.com/junk1tm/errorsx.svg)](https://pkg.go.dev/github.com/junk1tm/errorsx)
[![report](https://goreportcard.com/badge/github.com/junk1tm/errorsx)](https://goreportcard.com/report/github.com/junk1tm/errorsx)
[![codecov](https://codecov.io/gh/junk1tm/errorsx/branch/main/graph/badge.svg)](https://codecov.io/gh/junk1tm/errorsx)

Extensions for the standard `errors` package

## 📦 Install

```shell
go get github.com/junk1tm/errorsx
```

## 🧩 Extensions

### IsAny

A multi-target version of `errors.Is`.

Instead of:

```go
if errors.Is(err, os.ErrNotExist) || errors.Is(err, os.ErrPermission) {
	// handle error
}
```

Use this:

```go
if errorsx.IsAny(err, os.ErrNotExist, os.ErrPermission) {
	// handle error
}
```

### AsAny

A multi-target version of `errors.As`.

Instead of:

```go
if errors.As(err, new(*os.PathError)) || errors.As(err, new(*os.LinkError)) {
	// handle error
}
```

Use this:

```go
if errorsx.AsAny(err, new(*os.PathError), new(*os.LinkError)) {
	// handle error
}
```

### IsTimeout

Reports whether the error was caused by timeout. Unlike `os.IsTimeout`, it
respects error wrapping.

Won't catch a wrapped error:

```go
if os.IsTimeout(err) {
	// handle timeout
}
```

Will do just fine:

```go
if errorsx.IsTimeout(err) {
	// handle timeout
}
```

### Close

Attempts to close the given `io.Closer` and assigns the returned error (if any)
to `err`. Designed to be used ONLY as a defer statement.

```go
f, err := os.Open("file.txt")
if err != nil {
	return err
}
defer errorsx.Close(&err, f) // OR errorsx.Close(&err, f, "closing file: %w")
```
