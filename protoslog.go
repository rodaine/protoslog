// Package protoslog provides utilities for using protocol buffer messages with
// the log/slog package.
package protoslog

import (
	"log/slog"

	"google.golang.org/protobuf/proto"
)

// Message returns a [slog.Attr] for a [proto.Message] using the provided
// options. Note that these options are ignored if the [slog.Value] is handled
// by a [Handler].
func Message(key string, msg proto.Message, options ...Option) slog.Attr {
	return slog.Any(key, MessageValuer(msg, options...))
}

// MessageValue returns a [slog.Value] for a [proto.Message] using the provided
// options. Note that these options are ignored if the [slog.Value] is handled
// by a [Handler].
func MessageValue(msg proto.Message, options ...Option) slog.Value {
	return slog.AnyValue(MessageValuer(msg, options...))
}

// Valuer implements [slog.LogValuer] for a [proto.Message] to defer computing a
// [slog.Value] until it's needed.
type Valuer struct {
	// Message is the proto.Message to produce the slog.Value.
	Message proto.Message

	options options
}

// MessageValuer returns a [Valuer] for a [proto.Message] using the provided
// options. Note that these options are ignored if the [Valuer] is handled by a
// [Handler].
func MessageValuer(msg proto.Message, options ...Option) Valuer {
	return Valuer{Message: msg, options: newOptions(options)}
}

// LogValue satisfies the [slog.LogValuer] interface.
func (v Valuer) LogValue() slog.Value {
	if v.Message == nil {
		return slog.Value{}
	}
	return v.options.MessageValue(v.Message.ProtoReflect())
}
