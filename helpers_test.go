package protoslog_test

import (
	"log/slog"
	"os"
	"testing"

	"github.com/rodaine/protoslog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func slogHandler() slog.Handler {
	return slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if len(groups) == 0 && a.Key == "time" {
				return slog.Attr{}
			}
			return a
		},
	})
}

func testMessageValue(t *testing.T, name string, msg proto.Message, expected slog.Value, opts ...protoslog.Option) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		t.Parallel()
		val := protoslog.MessageValue(msg, opts...)
		assertSlogValue(t, expected, val)
	})
}

func assertSlogValue(t *testing.T, expected, actual slog.Value) {
	t.Helper()
	assert.True(t, actual.Resolve().Equal(expected.Resolve()),
		"expected %v, got %v", expected, actual)
}

func makeAny(t *testing.T, msg proto.Message) *anypb.Any {
	t.Helper()
	anyMsg, err := anypb.New(msg)
	require.NoError(t, err)
	return anyMsg
}
