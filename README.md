# protoslog [![Go Reference](https://pkg.go.dev/badge/github.com/rodaine/protoslog.svg)](https://pkg.go.dev/github.com/rodaine/protoslog) [![CI](https://github.com/rodaine/protoslog/actions/workflows/ci.yaml/badge.svg)](https://github.com/rodaine/protoslog/actions/workflows/ci.yaml)

`protoslog` provides utilities for using protocol buffer messages with the `log/slog` package introduced in Go 1.21.

## Example

`protoslog` operates against Protocol Buffer messages. Below, one might have such a `User` message:

```protobuf
syntax="proto3";

import "google/protobuf/timestamp.proto";

message User {
  fixed64 id = 1;
  string email = 2 [debug_redact=true];
  Status status = 3;
  google.protobuf.Timestamp updated = 4;
}

enum Status {
  UNSPECIFIED = 0;
  ACTIVE = 1;
  INACTIVE = 2;
}
```

`protoslog` does **NOT** require any code generation (beyond the output of `protoc-gen-go`) to properly log a message:

```go
package main

import (
	"log/slog"

	"github.com/rodaine/protoslog"
	"github.com/rodaine/protoslog/internal/gen"
)

func main() {
	msg := &gen.User{
		Id:      123,
		Email:   "rodaine@github.com",
		Status:  gen.ACTIVE,
		Updated: time.Now(),
	}

	slog.Info("hello", protoslog.Message("user", msg))
}
```

Outputs:

```text
2022/11/08 15:28:26 INFO hello user.id=123 user.email=REDACTED user.status=ACTIVE user.updated=2022-11-08T15:28:26.000Z
```

## Field Value Types

Messages are lazily converted into a `slog.GroupValue` with each of its populated field converted into a `slog.Attr` with the field name as the key and value produced based on its type (similar to the canonical JSON encoding rules)

### Scalar Types

- **bool**: `slog.BoolValue`
- **floats**: `slog.Float64Value`
- **bytes**: base64 encoded in a `slog.StringValue`
- **string**: `slog.StringValue`
- **enum**: `slog.StringValue` of the value name if it's defined, or `slog.Int64Value` otherwise
- **signed integer**: `slog.Int64Value`
- **unsigned integer**: `slog.Uint64Value`

### Composite Types

Populated composite fields are encoded as a `slog.GroupValue`:

- **message**: each field converted into a `slog.Attr` with its name as the key and the value recursively applying these rules
- **repeated**: each item converted into a `slog.Attr` with its index string-ified as the key and the value recursively applying these rules
- **map**: each entry converted into a `slog.Attr` with its key string-ified and the value recursively applying these rules

### Well-Known Types (WKTs)

Similar to the canonical JSON encoding, some of the WKTs produce special-cased `slog.Value`:

- **google.protobuf.NullValue**: empty `slog.Value{}` (equivalent of `nil`)\
- **google.protobuf.Timestamp**: `slog.TimeValue`
- **google.protobuf.Duration**: `slog.DurationValue`
- **wrappers**: it's `value` field, applying these rules
- **google.protobuf.ListValue**: its `values` field, applying the repeated rule above
- **google.protobuf.Struct**: its `fields` field, applying the map rule above
- **google.protobuf.Value**: the field set in its `kind` oneof, applying these rules
- **google.protobuf.Any**: see [Any WKT Resolution] below

## Redaction

Messages may contain personal identifiable information (PII), secrets, or
similar data that should not be written into a log. Message fields can be
annotated with the [debug_redact] option to identify such values. By default,
protoslog will redact these fields, with the behavior customizable via options.

Populated redacted fields are replaced with a `slog.StringValue("REDACTED")`:

```go
msg := &gen.User{Email: "personal@identifiable.info"}
slog.Info("default", protoslog.Message("user", msg))
// Stderr: 2022/11/08 15:28:26 INFO default user.email=REDACTED
```

To elide redacted fields instead of including them, `WithElideRedactions` can 
be used:

```go
slog.Info("elide", protoslog.Message("user", msg, protoslog.WithElideRedactions()))
// Stderr: 2022/11/08 15:28:26 INFO elide
```

Redaction may also be disabled via `WithDisableRedactions`:

```go
slog.Info("disable", protoslog.Message("user", msg, protoslog.WithDisableRedactions()))
// Stderr: 2022/11/08 15:28:26 INFO disable email=personal@identifiable.info
```

## All Fields

By default, `protoslog` only emits fields that are populated on the message (via
the behavior of `protoreflect.Message#Has`):

```go
msg := &gen.Location{Latitude: 1.23}
slog.Info("default", protoslog.Message("loc", msg))
// Stderr: 2022/11/08 15:28:26 INFO default loc.latitude=1.23
```

To emit all fields regardless of presence, use `WithAllFields`:

```go
slog.Info("all", protoslog.Message("loc", msg, protoslog.WithAllFields()))
// Stderr: 2022/11/08 15:28:26 INFO all loc.latitude=1.23 loc.longitude=0
```

For unpopulated "nullable," repeated, and map fields, the zero `slog.Value`
is emitted (which is equivalent to `nil`). All other fields emit their default 
values.

## Any WKT Resolution

`protoslog` emits the `Any` field's `type_url` with the key `@type`. By default,
`protoslog` attempts to resolve the field's value and on success emits it:

```go
msg := &gen.User{Id: 123}
anyPB, _ := anypb.New(msg)
slog.Info("success", protoslog.Message("any", anyPB))
// Stderr: 2022/11/08 15:28:26 INFO success any.@type=type.googleapis.com/User any.id=123
```

If the inner value does not resolve to a `slog.GroupValue` (e.g., it's a WKT), the result is added as `@value`:

```go
msg := durationpb.New(5*time.Second)
anyPB, _ := anypb.New(msg)
slog.Info("wkt", protoslog.Message("any", anyPB))
// Stderr: 2022/11/08 15:28:26 INFO wkt any.@type=type.googleapis.com/google.protobuf.Duration any.@value=5s
```

If the value cannot be resolved (either unknown or an error occurs), only the `@type` attribute will be present:

```go
anyPB := &anypb.Any{TypeUrl: "foobar"}
slog.Info("unknown", protoslog.Message("any", anyPB))
// Stderr: 2022/11/08 15:28:26 INFO unknown any.@type=foobar
```

By default, `protoslog` uses `protoregistry.GlobalTypes` to resolve Any WKTs. A custom resolver can be provided via `WithAnyResolver`:

```go
slog.Info("custom", protoslog.Message("any", anyPB, protoslog.WithAnyResolver(myResolver)))
```

To skip resolving Any WKTs, use `WithSkipAnys`. Only the `@type` attribute will be emitted:

```go
slog.Info("skip", protoslog.Message("any", anyPB, protoslog.WithSkipAnys()))
```

## `slog` Handler

If a message is not wrapped via `protoslog`, it will be presented in the logs 
with the behavior of `slog.AnyValue`. To ensure all messages are resolved 
correctly regardless, a `protoslog.Handler` can wrap a `slog.Handler`:

```go
handler := protoslog.NewHandler(slog.Default().Handler())
logger := slog.New(handler)

msg := &gen.User{Id: 123}
logger.Info("handler", "user", msg)
// Stderr: 2022/11/08 15:28:26 INFO handler user.id=123
```

The options on `protoslog.Handler` supersede those on messages wrapped via other 
`protoslog` functions.

## `protoc-gen-slog`

To make the generated message types produced by `protoc-gen-go` implement 
`slog.LogValuer`, `protoc-gen-slog` can be used to generate `LogValue` methods.

```shell
go install github.com/rodaine/protoslog/protoc-gen-slog
```

### Buf CLI

When using `buf`, ensure the `out` path and `opt` values are equivalent for both
`protoc-gen-go` and `protoc-gen-slog` plugins:

```yaml
# buf.gen.yaml
version: v1
plugins:
  - plugin: buf.build/protocolbuffers/go:v1.32.0
    out: gen
    opt:
      - paths=source_relative
  - plugin: slog
    out: gen
    opt:
      - paths=source_relative
```

### protoc

When using `protoc`, ensure both plugin options and output path are equivalent:

```shell
protoc \
  --go_out="$OUT" \
  --slog_out="$OUT" \
  $PROTOS
```

[debug_redact]: https://github.com/protocolbuffers/protobuf/blob/v22.0/src/google/protobuf/descriptor.proto#L630-L632