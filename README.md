# idf

ID formatting library. Use it to format existing values as IDs in specific format (e.g. normalized, padded, predictable, compounded, etc.).

[![Test](https://github.com/mchmarny/idf/actions/workflows/test-on-push.yaml/badge.svg?branch=main)](https://github.com/mchmarny/idf/actions/workflows/test-on-push.yaml) [![Go Report Card](https://goreportcard.com/badge/github.com/mchmarny/idf)](https://goreportcard.com/report/github.com/mchmarny/idf) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mchmarny/idf) [![codecov](https://codecov.io/gh/mchmarny/idf/branch/main/graph/badge.svg?token=00H8S7GMPP)](https://codecov.io/gh/mchmarny/idf)

> Note, this library is under development, don't use in anything serious just yet. 

## Usage

The `idf` library can be configured with any number of the supported formatters (see the supported list of formatters below). 

For example, import:

```go
import "github.com/mchmarny/idf/id"
```

and than create an instance: 

```go
f := id.New(id.WithPadding("^", 20),
		    id.WithSHA256Encoding(),
		    id.WithPrefix("id-"),
		    id.WithSuffix("-test"),
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
f := id.New(id.WithBase64Encoding(), id.WithPrefix("id-"))
id, err := f.ToID("123-^!(")
```

Will result in `id` being `id-MTIzNDU2Nzg5MTAxMTEyMTMxNDE1MTYxNzE4MTkyMDIxMjIyMzI0MjUyNi1eISg=`

### Predictable Length ID

You want values of different lengths to result in a consistent length ID.

```go
f := id.New(id.WithSHA256Encoding(), id.WithPrefix("id-"))
id1, err := f.ToID("user-12345")
id2, err := f.ToID("user-12345678910111213141516171819")
```

Will result in `id1` being:
`id-43a2f41a7bffacce74013d74a2f459db5d69d38e061cb8d0e5e262102e2d98d7` 
and `id2` being: 
`id-e99f84f908862e2ea62f7ee4db4af40322a0adf9174f085ec64e032fadbbff56`

### Predictable and Compounded ID

You need to create reproducible keys from multiple values which allow for reliable scans/sorts (e.g. key in a NoSQL systems without secondary index support).

```go
f := id.New(id.WithSHA256Encoding(), 
             id.WithPrefix("id-"), 
			 id.WithDatetime("2006-01-02", "-", true),
	)
id1, err := f.ToID("user-12345")
```

Will result in `id1` being: `id-43a2f41a7bffacce74013d74a2f459db5d69d38e061cb8d0e5e262102e2d98d7-2022-01-01`, which later would allow you to reliably query on that compounded key. 

```sql
where key >= 'id-43a2f41a7bffacce74013d74a2f459db5d69d38e061cb8d0e5e262102e2d98d7-2022-01-01' and key <= 'id-43a2f41a7bffacce74013d74a2f459db5d69d38e061cb8d0e5e262102e2d98d7-2022-01-31'
```