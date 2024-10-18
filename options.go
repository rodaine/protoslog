package protoslog

import (
	"cmp"
	"encoding/base64"
	"log/slog"
	"slices"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

//nolint:gochecknoglobals // these are effectively constants
var (
	redactedValue = slog.StringValue("REDACTED")
	unknownValue  = slog.StringValue("UNKNOWN")
)

// Option functions customize generation of a [slog.Value] from a
// [proto.Message] beyond the default behavior.
type Option func(o *options)

// WithDisableRedactions indicates that fields annotated with the debug_redact
// option should not be redacted from the [slog.Value].
func WithDisableRedactions() Option {
	return func(o *options) {
		o.DisableRedaction = true
	}
}

// WithElideRedactions indicates that redacted fields should be removed from the
// [slog.Value] instead of being replaced with REDACTED. WithElideRedactions
// supersedes [WithAllFields].
func WithElideRedactions() Option {
	return func(o *options) { o.ElideRedacted = true }
}

// WithAllFields indicates that all fields, including unpopulated ones, should be
// included in [slog.Value]. Unpopulated members of a oneof are still excluded
// from the output.
func WithAllFields() Option {
	return func(o *options) { o.IncludeAllFields = true }
}

// WithSkipAnys indicates that google.protobuf.Any fields should not be
// unmarshalled during construction of the [slog.Value], emitting only a "@type"
// field.
func WithSkipAnys() Option {
	return func(o *options) { o.SkipAnyResolution = true }
}

// WithAnyResolver is the [protoregistry.MessageTypeResolver] used to resolve
// the google.protobuf.Any well-known type into a valid [slog.Value]. When nil,
// [protoregistry.GlobalTypes] is used. If the type cannot be found in the
// resolver or if unmarshaling fails, only a "@type" field is emitted.
func WithAnyResolver(resolver protoregistry.MessageTypeResolver) Option {
	return func(o *options) { o.AnyResolver = resolver }
}

type options struct {
	// DisableRedaction indicates that fields annotated with the debug_redact
	// option should not be redacted from the slog.Value. When false, redacted
	// field values are not emitted.
	DisableRedaction bool
	// ElideRedacted indicates that redacted fields should be removed from the
	// slog.Value. When false, redacted fields are replaced with a "REDACTED"
	// string. ElideRedacted supersedes IncludeAllFields.
	ElideRedacted bool
	// IncludeAllFields indicates that all fields, including unpopulated ones, should be
	// included in slog.Value. Unpopulated members of a oneof are still excluded
	// from the output. When false, only populated fields are included in the output.
	IncludeAllFields bool
	// SkipAnyResolution indicates that google.protobuf.Any fields should not be
	// unmarshaled during construction of the slog.Value, emitting only an "@type"
	// field. When false, Any fields will attempt to be unmarshaled using the
	// specified AnyResolver.
	SkipAnyResolution bool
	// AnyResolver is the protoregistry.MessageTypeResolver used to resolve
	// the google.protobuf.Any well-known type into a valid slog.Value. When nil,
	// protoregistry.GlobalTypes is used. If the type cannot be found in the
	// resolver or if unmarshaling fails, only an "@type" field is emitted.
	AnyResolver protoregistry.MessageTypeResolver
}

func newOptions(opts []Option) (o options) {
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

func mergeOptions(base, other options) (o options) {
	o.DisableRedaction = base.DisableRedaction || other.DisableRedaction
	o.ElideRedacted = base.ElideRedacted || other.ElideRedacted
	o.IncludeAllFields = base.IncludeAllFields || other.IncludeAllFields
	o.SkipAnyResolution = base.SkipAnyResolution || other.SkipAnyResolution
	o.AnyResolver = base.AnyResolver
	if other.AnyResolver != nil {
		o.AnyResolver = other.AnyResolver
	}
	return o
}

// MessageValue produces a [slog.Value] for the provided [protoreflect.Message].
func (o options) MessageValue(msg protoreflect.Message) slog.Value {
	if msg == nil || !msg.IsValid() {
		return slog.Value{}
	}
	if val, ok := o.wktValue(msg); ok {
		return val
	}
	var attrs []slog.Attr
	fields := msg.Descriptor().Fields()
	for i, n := 0, fields.Len(); i < n; i++ {
		field := fields.Get(i)
		fieldOpts, _ := field.Options().(*descriptorpb.FieldOptions)
		var val slog.Value
		switch {
		case !o.IncludeAllFields && !msg.Has(field):
			// unpopulated field
			continue
		case field.ContainingOneof() != nil && !msg.Has(field):
			// unpopulated oneof field
			continue
		case !o.DisableRedaction && fieldOpts.GetDebugRedact():
			// redacted field
			if o.ElideRedacted {
				continue
			}
			val = redactedValue
		case field.IsList():
			val = o.listValue(field, msg.Get(field).List())
		case field.IsMap():
			val = o.mapValue(field, msg.Get(field).Map())
		default:
			val = o.singularValue(field, msg.Get(field))
		}
		attrs = append(attrs, slog.Attr{Key: string(field.Name()), Value: val})
	}
	return slog.GroupValue(attrs...)
}

func (o options) listValue(field protoreflect.FieldDescriptor, value protoreflect.List) slog.Value {
	length := value.Len()
	if length == 0 {
		return slog.Value{}
	}
	attrs := make([]slog.Attr, length)
	for i := range length {
		item := value.Get(i)
		attrs[i].Key = strconv.Itoa(i)
		attrs[i].Value = o.singularValue(field, item)
	}
	return slog.GroupValue(attrs...)
}

func (o options) mapValue(field protoreflect.FieldDescriptor, value protoreflect.Map) slog.Value {
	length := value.Len()
	if length == 0 {
		return slog.Value{}
	}
	var cmpFunc func(a, b protoreflect.MapKey) int
	//nolint:exhaustive // map key kinds are a subset of all protoreflect.Kind values
	switch field.MapKey().Kind() {
	case protoreflect.BoolKind:
		cmpFunc = cmpBool
	case protoreflect.Int32Kind, protoreflect.Int64Kind,
		protoreflect.Sint32Kind, protoreflect.Sint64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
		cmpFunc = cmpInt
	case protoreflect.Uint32Kind, protoreflect.Uint64Kind,
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
		cmpFunc = cmpUint
	case protoreflect.StringKind:
		cmpFunc = cmpString
	default:
		return unknownValue
	}
	keys := make([]protoreflect.MapKey, 0, length)
	attrs := make([]slog.Attr, length)
	value.Range(func(key protoreflect.MapKey, _ protoreflect.Value) bool {
		keys = append(keys, key)
		return true
	})
	slices.SortFunc(keys, cmpFunc)
	valDesc := field.MapValue()
	for i, key := range keys {
		attrs[i].Key = key.String()
		attrs[i].Value = o.singularValue(valDesc, value.Get(key))
	}
	return slog.GroupValue(attrs...)
}

func (o options) singularValue(field protoreflect.FieldDescriptor, value protoreflect.Value) slog.Value {
	switch field.Kind() {
	case protoreflect.BoolKind:
		return slog.BoolValue(value.Bool())
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return slog.Float64Value(value.Float())
	case protoreflect.BytesKind:
		return slog.StringValue(base64.StdEncoding.EncodeToString(value.Bytes()))
	case protoreflect.StringKind:
		return slog.StringValue(value.String())
	case protoreflect.EnumKind:
		eDesc := field.Enum()
		if eDesc.FullName() == "google.protobuf.NullValue" {
			return slog.Value{}
		}
		evDesc := eDesc.Values().ByNumber(value.Enum())
		if evDesc == nil {
			return slog.Int64Value(int64(value.Enum()))
		}
		return slog.StringValue(string(evDesc.Name()))
	case protoreflect.Int32Kind, protoreflect.Int64Kind,
		protoreflect.Sint32Kind, protoreflect.Sint64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
		return slog.Int64Value(value.Int())
	case protoreflect.Uint32Kind, protoreflect.Uint64Kind,
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
		return slog.Uint64Value(value.Uint())
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return o.MessageValue(value.Message())
	default:
		return unknownValue
	}
}

func (o options) wktValue(msg protoreflect.Message) (slog.Value, bool) {
	switch msg.Descriptor().FullName() {
	case "google.protobuf.Timestamp":
		secsDesc := msg.Descriptor().Fields().ByName("seconds")
		nanosDesc := msg.Descriptor().Fields().ByName("nanos")
		ts := time.Unix(msg.Get(secsDesc).Int(), msg.Get(nanosDesc).Int())
		return slog.TimeValue(ts), true
	case "google.protobuf.Duration":
		secsDesc := msg.Descriptor().Fields().ByName("seconds")
		nanosDesc := msg.Descriptor().Fields().ByName("nanos")
		secs := time.Duration(msg.Get(secsDesc).Int()) * time.Second
		nanos := time.Duration(msg.Get(nanosDesc).Int()) * time.Nanosecond
		return slog.DurationValue(secs + nanos), true
	case "google.protobuf.Any":
		typeDesc := msg.Descriptor().Fields().ByName("type_url")
		typeURL := msg.Get(typeDesc).String()
		attrs := []slog.Attr{{
			Key:   "@type",
			Value: slog.StringValue(typeURL),
		}}
		if !o.SkipAnyResolution {
			attrs = append(attrs, o.resolveAnyValue(msg, typeURL)...)
		}
		return slog.GroupValue(attrs...), true
	case "google.protobuf.BoolValue",
		"google.protobuf.BytesValue",
		"google.protobuf.DoubleValue",
		"google.protobuf.FloatValue",
		"google.protobuf.Int32Value",
		"google.protobuf.Int64Value",
		"google.protobuf.StringValue",
		"google.protobuf.UInt32Value",
		"google.protobuf.UInt64Value":
		valueDesc := msg.Descriptor().Fields().ByName("value")
		return o.singularValue(valueDesc, msg.Get(valueDesc)), true
	case "google.protobuf.ListValue":
		valuesDesc := msg.Descriptor().Fields().ByName("values")
		return o.listValue(valuesDesc, msg.Get(valuesDesc).List()), true
	case "google.protobuf.Value":
		kindDesc := msg.Descriptor().Oneofs().ByName("kind")
		fldDesc := msg.WhichOneof(kindDesc)
		if fldDesc == nil {
			// considered an error to not be set
			return unknownValue, true
		}
		return o.singularValue(fldDesc, msg.Get(fldDesc)), true
	case "google.protobuf.Struct":
		fieldsDesc := msg.Descriptor().Fields().ByName("fields")
		return o.mapValue(fieldsDesc, msg.Get(fieldsDesc).Map()), true
	default:
		return slog.Value{}, false
	}
}

func (o options) resolveAnyValue(msg protoreflect.Message, typeURL string) []slog.Attr {
	resolver := o.AnyResolver
	if resolver == nil {
		resolver = protoregistry.GlobalTypes
	}
	msgTyp, err := resolver.FindMessageByURL(typeURL)
	if err != nil {
		return nil
	}

	anyMsg := msgTyp.New()
	valueDesc := msg.Descriptor().Fields().ByName("value")
	err = proto.Unmarshal(msg.Get(valueDesc).Bytes(), anyMsg.Interface())
	if err != nil {
		return nil
	}

	anyVal := o.MessageValue(anyMsg)
	if anyVal.Kind() == slog.KindGroup {
		return anyVal.Group()
	}
	return []slog.Attr{slog.Any("@value", anyVal)}
}

func cmpBool(a, b protoreflect.MapKey) int {
	switch x, y := a.Bool(), b.Bool(); {
	case !x && y:
		return -1
	case x == y:
		return 0
	default:
		return 1
	}
}

func cmpInt(a, b protoreflect.MapKey) int {
	return cmp.Compare(a.Int(), b.Int())
}

func cmpUint(a, b protoreflect.MapKey) int {
	return cmp.Compare(a.Uint(), b.Uint())
}

func cmpString(a, b protoreflect.MapKey) int {
	return cmp.Compare(a.String(), b.String())
}
