package protoslog_test

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/rodaine/protoslog"
	pb "github.com/rodaine/protoslog/internal/gen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestMain(m *testing.M) {
	time.Local = time.UTC
	os.Exit(m.Run())
}

func ExampleMessage() {
	updated := time.Date(2012, time.September, 2, 15, 53, 0, 0, time.UTC)
	msg := &pb.User{
		Id:    123,
		Name:  "foobar",
		Email: "foo@bar.com", // debug_redact
		Location: &pb.Location{
			Latitude:  1.23,
			Longitude: 4.56,
		},
		Hobbies: []string{"track", "field"},
		Pets: map[string]pb.PetType{
			"Rover": pb.PetType_PET_TYPE_DOG,
			"Fifi":  pb.PetType_PET_TYPE_CAT,
		},
		Updated:       timestamppb.New(updated),
		Best_100MTime: durationpb.New(9*time.Second + 580*time.Millisecond),
	}

	logger := slog.New(slogHandler())
	logger.Info("some event", protoslog.Message("user", msg))
	// Output:
	// level=INFO msg="some event" user.id=123 user.name=foobar user.email=REDACTED user.location.latitude=1.23 user.location.longitude=4.56 user.hobbies.0=track user.hobbies.1=field user.pets.Fifi=PET_TYPE_CAT user.pets.Rover=PET_TYPE_DOG user.updated=2012-09-02T15:53:00.000Z user.best_100m_time=9.58s
}

func TestMessage(t *testing.T) {
	t.Parallel()
	attr := protoslog.Message("foo", &pb.Singulars{Bool: true})
	assert.Equal(t, "foo", attr.Key)
	assertSlogValue(t,
		slog.GroupValue(slog.Bool("bool", true)),
		attr.Value,
	)
}

func TestMessageValue(t *testing.T) {
	t.Parallel()

	testMessageValue(t, "nil", nil, slog.Value{})

	t.Run("defaults", func(t *testing.T) {
		t.Parallel()
		testMessageValue(t, "empty_singulars", &pb.Singulars{}, slog.GroupValue())
		testMessageValue(t, "all_singulars", &pb.Singulars{
			Bool:     true,
			Float:    2.5,
			Double:   4.56,
			Bytes:    []byte{'a', 'b', 'c'},
			String_:  "foobar",
			Enum:     pb.Enum_ENUM_ONE,
			Int32:    789,
			Int64:    1011,
			Sint32:   1213,
			Sint64:   1415,
			Sfixed32: 1617,
			Sfixed64: 1819,
			Uint32:   2021,
			Uint64:   2223,
			Fixed32:  2425,
			Fixed64:  2627,
			Message:  &pb.Singulars{Bool: true},
		}, slog.GroupValue(
			slog.Bool("bool", true),
			slog.Float64("float", 2.5),
			slog.Float64("double", 4.56),
			slog.String("bytes", "YWJj"),
			slog.String("string", "foobar"),
			slog.String("enum", "ENUM_ONE"),
			slog.Int64("int32", 789),
			slog.Int64("int64", 1011),
			slog.Int64("sint32", 1213),
			slog.Int64("sint64", 1415),
			slog.Int64("sfixed32", 1617),
			slog.Int64("sfixed64", 1819),
			slog.Uint64("uint32", 2021),
			slog.Uint64("uint64", 2223),
			slog.Uint64("fixed32", 2425),
			slog.Uint64("fixed64", 2627),
			slog.Group("message", slog.Bool("bool", true)),
		))
		testMessageValue(t, "oneof_empty", &pb.Oneof{}, slog.GroupValue())
		testMessageValue(t, "oneof_set", &pb.Oneof{O: &pb.Oneof_Int32{Int32: 123}}, slog.GroupValue(slog.Int64("int32", 123)))
		testMessageValue(t, "lists_empty", &pb.Lists{}, slog.GroupValue())
		testMessageValue(t, "lists_set", &pb.Lists{
			Singulars: []int32{1, 2, 3},
			Messages:  []*pb.Singulars{{Bool: true}, {Int32: 123}},
		}, slog.GroupValue(
			slog.Group("singulars",
				slog.Int64("0", 1),
				slog.Int64("1", 2),
				slog.Int64("2", 3),
			),
			slog.Group("messages",
				slog.Group("0", slog.Bool("bool", true)),
				slog.Group("1", slog.Int64("int32", 123)),
			),
		))
		testMessageValue(t, "maps_empty", &pb.Maps{}, slog.GroupValue())
		testMessageValue(t, "maps_set", &pb.Maps{
			Bools:   map[bool]string{true: "true"},
			Ints:    map[int32]string{123: "123"},
			Uints:   map[uint32]string{789: "789"},
			Strings: map[string]string{"foo": "bar"},
		}, slog.GroupValue(
			slog.Group("bools", slog.String("true", "true")),
			slog.Group("ints", slog.String("123", "123")),
			slog.Group("uints", slog.String("789", "789")),
			slog.Group("strings", slog.String("foo", "bar")),
		))
	})

	t.Run("all_fields", func(t *testing.T) {
		t.Parallel()

		opts := []protoslog.Option{protoslog.WithAllFields()}

		testMessageValue(t, "singulars", &pb.Singulars{}, slog.GroupValue(
			slog.Bool("bool", false),
			slog.Float64("float", 0),
			slog.Float64("double", 0),
			slog.String("bytes", ""),
			slog.String("string", ""),
			slog.String("enum", "ENUM_UNSPECIFIED"),
			slog.Int64("int32", 0),
			slog.Int64("int64", 0),
			slog.Int64("sint32", 0),
			slog.Int64("sint64", 0),
			slog.Int64("sfixed32", 0),
			slog.Int64("sfixed64", 0),
			slog.Uint64("uint32", 0),
			slog.Uint64("uint64", 0),
			slog.Uint64("fixed32", 0),
			slog.Uint64("fixed64", 0),
			slog.Attr{Key: "message", Value: slog.Value{}},
		), opts...)

		testMessageValue(t, "oneofs", &pb.Oneof{}, slog.GroupValue(), opts...)

		testMessageValue(t, "lists", &pb.Lists{}, slog.GroupValue(
			slog.Attr{Key: "singulars", Value: slog.Value{}},
			slog.Attr{Key: "messages", Value: slog.Value{}},
		), opts...)

		testMessageValue(t, "maps", &pb.Maps{}, slog.GroupValue(
			slog.Attr{Key: "bools", Value: slog.Value{}},
			slog.Attr{Key: "ints", Value: slog.Value{}},
			slog.Attr{Key: "uints", Value: slog.Value{}},
			slog.Attr{Key: "strings", Value: slog.Value{}},
		), opts...)
	})

	t.Run("redaction", func(t *testing.T) {
		t.Parallel()

		testMessageValue(t, "default", &pb.Redaction{Val: 123}, slog.GroupValue(slog.Attr{Key: "val", Value: slog.StringValue("REDACTED")}))

		testMessageValue(t, "disable_redaction", &pb.Redaction{Val: 123}, slog.GroupValue(slog.Int64("val", 123)), protoslog.WithDisableRedactions())

		testMessageValue(t, "elide_redaction", &pb.Redaction{Val: 123}, slog.GroupValue(), protoslog.WithElideRedactions())

		testMessageValue(t, "all_fields", &pb.Redaction{}, slog.GroupValue(slog.Attr{Key: "val", Value: slog.StringValue("REDACTED")}), protoslog.WithAllFields())

		testMessageValue(t, "all_fields_elide", &pb.Redaction{}, slog.GroupValue(), protoslog.WithElideRedactions(), protoslog.WithAllFields())
	})

	t.Run("wkts", func(t *testing.T) {
		t.Parallel()

		now := timestamppb.Now()
		testMessageValue(t, "timestamp", now, slog.TimeValue(now.AsTime()))

		dur := durationpb.New(5 * time.Minute)
		testMessageValue(t, "duration", dur, slog.DurationValue(dur.AsDuration()))

		anyMsg := makeAny(t, &pb.Singulars{Bool: true})
		testMessageValue(t, "any", anyMsg, slog.GroupValue(
			slog.String("@type", anyMsg.GetTypeUrl()),
			slog.Bool("bool", true),
		))

		anyWkt := makeAny(t, now)
		testMessageValue(t, "any_wkt", anyWkt, slog.GroupValue(
			slog.String("@type", anyWkt.GetTypeUrl()),
			slog.Time("@value", now.AsTime()),
		))

		testMessageValue(t, "bool_value", wrapperspb.Bool(true), slog.BoolValue(true))
		testMessageValue(t, "bytes_value", wrapperspb.Bytes([]byte{'a', 'b', 'c'}), slog.StringValue("YWJj"))
		testMessageValue(t, "double_value", wrapperspb.Double(1.23), slog.Float64Value(1.23))
		testMessageValue(t, "float_value", wrapperspb.Float(4.5), slog.Float64Value(4.5))
		testMessageValue(t, "int32_value", wrapperspb.Int32(123), slog.Int64Value(123))
		testMessageValue(t, "int64_value", wrapperspb.Int64(456), slog.Int64Value(456))
		testMessageValue(t, "string_value", wrapperspb.String("foo"), slog.StringValue("foo"))
		testMessageValue(t, "uint32_value", wrapperspb.UInt32(123), slog.Uint64Value(123))
		testMessageValue(t, "uint64_value", wrapperspb.UInt64(456), slog.Uint64Value(456))

		testMessageValue(t, "null", structpb.NewNullValue(), slog.Value{})

		list, err := structpb.NewList([]any{"foo", true, nil})
		require.NoError(t, err)
		testMessageValue(t, "list", list, slog.GroupValue(
			slog.String("0", "foo"),
			slog.Bool("1", true),
			slog.Attr{Key: "2"},
		))

		str, err := structpb.NewStruct(map[string]any{
			"foo":  true,
			"bar":  "baz",
			"quux": nil,
		})
		require.NoError(t, err)
		testMessageValue(t, "struct", str, slog.GroupValue(
			slog.String("bar", "baz"),
			slog.Bool("foo", true),
			slog.Attr{Key: "quux"},
		))
	})

	t.Run("key-formatter", func(t *testing.T) {
		t.Parallel()

		keyFormatterFn := func(key string, fd protoreflect.FieldDescriptor, val protoreflect.Value) string {
			if fd.Kind() != protoreflect.MessageKind {
				return key
			}

			if val.Message().Descriptor().FullName() != "google.protobuf.Any" {
				return key
			}

			typeDesc := val.Message().Descriptor().Fields().ByName("type_url")
			typeURL := val.Message().Get(typeDesc).String()
			return fmt.Sprintf(
				"%s[%s]",
				key,
				strings.TrimPrefix(typeURL, "type.googleapis.com/"),
			)
		}

		payloadA := makeAny(t, &pb.PayloadA{Id: "some-string-based-id"})
		payloadB := makeAny(t, &pb.PayloadB{Id: uint64(123)})

		now := timestamppb.Now()

		testMessageValue(t, "with_formatter", &pb.Event{Payload: payloadA, Timestamp: now}, slog.GroupValue(
			slog.Attr{
				Key: "payload[PayloadA]",
				Value: slog.GroupValue(
					slog.String("@type", payloadA.GetTypeUrl()),
					slog.String("id", "some-string-based-id"),
				),
			},
			slog.Time("timestamp", now.AsTime()),
		), protoslog.WithKeyFormatter(keyFormatterFn))
		testMessageValue(t, "with_formatter", &pb.Event{Payload: payloadB, Timestamp: now}, slog.GroupValue(
			slog.Attr{
				Key: "payload[PayloadB]",
				Value: slog.GroupValue(
					slog.String("@type", payloadB.GetTypeUrl()),
					slog.Uint64("id", 123),
				),
			},
			slog.Time("timestamp", now.AsTime()),
		), protoslog.WithKeyFormatter(keyFormatterFn))

		testMessageValue(t, "without_formatter", &pb.Event{Payload: payloadA, Timestamp: now}, slog.GroupValue(
			slog.Attr{
				Key: "payload",
				Value: slog.GroupValue(
					slog.String("@type", payloadA.GetTypeUrl()),
					slog.String("id", "some-string-based-id"),
				),
			},
			slog.Time("timestamp", now.AsTime()),
		))
		testMessageValue(t, "without_formatter", &pb.Event{Payload: payloadB, Timestamp: now}, slog.GroupValue(
			slog.Attr{
				Key: "payload",
				Value: slog.GroupValue(
					slog.String("@type", payloadB.GetTypeUrl()),
					slog.Uint64("id", 123),
				),
			},
			slog.Time("timestamp", now.AsTime()),
		))
	})
}
