# idf

ID formatting library. Use it to format existing values as IDs in specific format (e.g. normalized, padded, predictable, compounded, etc.).

[![Test](https://github.com/mchmarny/idf/actions/workflows/test-on-push.yaml/badge.svg?branch=main)](https://github.com/mchmarny/idf/actions/workflows/test-on-push.yaml) [![Go Report Card](https://goreportcard.com/badge/github.com/mchmarny/idf)](https://goreportcard.com/report/github.com/mchmarny/idf) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mchmarny/idf) [![codecov](https://codecov.io/gh/mchmarny/idf/branch/main/graph/badge.svg?token=00H8S7GMPP)](https://codecov.io/gh/mchmarny/idf)

## Usage

The `idf` library can be configured with any number of the supported formatters (see the supported list of formatters below). For example:

```go
f := idf.New(
        idf.WithPadding("^", 20),
		idf.WithSHA256Encoding(),
		idf.WithPrefix("id-"),
		idf.WithSuffix("-test"),
    )
```

> Note, the order of formatters matters! Padding and encoding will render different results than encoding and then padding. 

Regardless of the type and number of formatters with which `idf` was configured, it provides a single `ToID(v string)` function which returns the formatted ID or an error. For example: 

```go
id, err := f.ToID("User1234")
```

## Formatters

The `idf` library currently supports the following formatters:

* `WithNormalizing()` - normalizes the input to lowercase and trims spaces.
* `WithPadding(char string, length int)` - pads the input at the end with the given character up to the given length.
* `WithBase64Encoding()` - base64 and URL encodes the input.
* `WithPrefix(prefix string)` - prepends the provided prefix to the input.
* `WithSuffix(suffix string)` - appends the provided suffix to the input.
* `WithSHA256Encoding()` - sha256 encodes the input.
* `WithDatetime(format, separator string, v time.Time)` - formats the time value and appends it to the input with the given separator. 

## Use-cases

> There is a few of fully functional code samples in the [examples](./examples) folder.

### Safe ID

Value you don't control has to be used as an ID in a system with specific ID format requirements. For example, IDs must be start with alpha and be URL safe. 

```go
f := idf.New(idf.WithBase64Encoding(), idf.WithPrefix("id-"))
id, err := f.ToID("123-^!(")
```

Will result in `id` being `id-MTIzNDU2Nzg5MTAxMTEyMTMxNDE1MTYxNzE4MTkyMDIxMjIyMzI0MjUyNi1eISg=`

### Predictable ID

WIP

### Predictably Compounded ID

WIP