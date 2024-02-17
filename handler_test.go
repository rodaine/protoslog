package protoslog_test

import (
	"context"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/rodaine/protoslog"
	pb "github.com/rodaine/protoslog/internal/gen"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
)

func TestHandler_Enabled(t *testing.T) {
	t.Parallel()

	handler := protoslog.NewHandler(slog.NewTextHandler(
		io.Discard,
		&slog.HandlerOptions{Level: slog.LevelInfo},
	))

	assert.True(t, handler.Enabled(context.Background(), slog.LevelInfo))
	assert.False(t, handler.Enabled(context.Background(), slog.LevelDebug))
}

func ExampleHandler() {
	msg := &pb.User{
		Id:            123,
		Best_100MTime: durationpb.New(9*time.Second + 580*time.Millisecond),
	}

	handler := protoslog.NewHandler(slogHandler())
	logger := slog.New(handler)
	logger.Info("hello world", "user", msg)
	// Output:
	// level=INFO msg="hello world" user.id=123 user.best_100m_time=9.58s
}

// Messages may contain personal identifiable information (PII), secrets, or
// similar data that should not be written into a log. Message fields can be
// annotated with the [debug_redact] option to identify such values. By default,
// protoslog will redact these fields, with the behavior customizable via
// [options].
//
// [debug_redact]: https://github.com/protocolbuffers/protobuf/blob/v22.0/src/google/protobuf/descriptor.proto#L630-L632
func ExampleHandler_redaction() {
	// message User { fixed64 id = 1; string email = 2 [debug_redact=true]; }
	msg := &pb.User{
		Id:    123,
		Email: "personal@identifiable.info",
	}

	childHandler := slogHandler()
	logger := slog.New(protoslog.NewHandler(childHandler))
	logger.Info("default", "user", msg)

	logger = slog.New(protoslog.NewHandler(childHandler, protoslog.WithDisableRedactions()))
	logger.Info("disabled", "user", msg)

	logger = slog.New(protoslog.NewHandler(childHandler, protoslog.WithElideRedactions()))
	logger.Info("elided", "user", msg)
	// Output:
	// level=INFO msg=default user.id=123 user.email=REDACTED
	// level=INFO msg=disabled user.id=123 user.email=personal@identifiable.info
	// level=INFO msg=elided user.id=123
}

func ExampleHandler_allFields() {
	msg := &pb.Location{
		Latitude: 1.23,
	}

	childHandler := slogHandler()
	logger := slog.New(protoslog.NewHandler(childHandler))
	logger.Info("default", "loc", msg)

	logger = slog.New(protoslog.NewHandler(childHandler, protoslog.WithAllFields()))
	logger.Info("all", "loc", msg)
	// Output:
	// level=INFO msg=default loc.latitude=1.23
	// level=INFO msg=all loc.latitude=1.23 loc.longitude=0
}

func ExampleHandler_any() {
	msg, _ := anypb.New(&pb.Location{
		Latitude:  1.23,
		Longitude: 4.56,
	})

	childHandler := slogHandler()
	logger := slog.New(protoslog.NewHandler(childHandler))
	logger.Info("default", "any", msg)
	logger.Info("unknown", "any", &anypb.Any{TypeUrl: "foobar"})

	logger = slog.New(protoslog.NewHandler(childHandler, protoslog.WithSkipAnys()))
	logger.Info("skip", "any", msg)
	// Output:
	// level=INFO msg=default any.@type=type.googleapis.com/Location any.latitude=1.23 any.longitude=4.56
	// level=INFO msg=unknown any.@type=foobar
	// level=INFO msg=skip any.@type=type.googleapis.com/Location
}

func ExampleHandler_WithAttrs() {
	loc := &pb.Location{Latitude: 1.23}
	msg := &pb.User{Id: 456}

	logger := slog.New(protoslog.NewHandler(slogHandler()))
	logger.With("loc", loc).Info("attrs", "user", msg)
	// Output:
	// level=INFO msg=attrs loc.latitude=1.23 user.id=456
}

func ExampleHandler_WithGroup() {
	msg := &pb.User{Id: 123}

	logger := slog.New(protoslog.NewHandler(slogHandler()))
	logger.WithGroup("foo").Info("group", "user", msg)
	// Output:
	// level=INFO msg=group foo.user.id=123
}
