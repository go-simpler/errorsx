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

### IsOneOf

A multi-target version of `errors.Is`.

### AsOneOf

A multi-target version of `errors.As`.

### IsTimeout

Reports whether the error was caused by timeout. Unlike `os.IsTimeout`, it
respects error wrapping.
