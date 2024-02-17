package protoslog

import (
	"context"
	"log/slog"

	"google.golang.org/protobuf/proto"
)

// Handler is a [slog.Handler] that properly converts [proto.Message] attributes
// into the appropriate [slog.Value], before delegating to a child [slog.Handler].
// The Handler's options merge with the options associated with [proto.Message]
// types that implement [slog.LogValuer]. Handler must be constructed via [NewHandler].
type Handler struct {
	child   slog.Handler
	options options
}

// NewHandler creates a [Handler] that delegates to child, using the given
// options. Note that these options merge with any options used in [Message],
// [MessageValue], or [MessageValuer].
func NewHandler(child slog.Handler, options ...Option) *Handler {
	return &Handler{child: child, options: newOptions(options)}
}

// Enabled delegates this check to its child handler.
func (h Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.child.Enabled(ctx, level)
}

// Handle converts the [proto.Message] attributes on record before delegating
// the record to its child handler.
func (h Handler) Handle(ctx context.Context, record slog.Record) error {
	attrs := make([]slog.Attr, 0, record.NumAttrs())
	record.Attrs(func(attr slog.Attr) bool {
		attrs = append(attrs, h.wrapAttr(attr))
		return true
	})
	record = slog.NewRecord(
		record.Time,
		record.Level,
		record.Message,
		record.PC)
	record.AddAttrs(attrs...)
	return h.child.Handle(ctx, record)
}

// WithAttrs converts the [proto.Message] attributes before delegating them to
// its child handler.
func (h Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h.wrapAttrs(attrs)
	return Handler{
		child:   h.child.WithAttrs(attrs),
		options: h.options,
	}
}

// WithGroup delegates starting this group on its child handler.
func (h Handler) WithGroup(name string) slog.Handler {
	return Handler{
		child:   h.child.WithGroup(name),
		options: h.options,
	}
}

func (h Handler) wrapAttrs(attrs []slog.Attr) {
	for i := range attrs {
		attrs[i] = h.wrapAttr(attrs[i])
	}
}

func (h Handler) wrapAttr(attr slog.Attr) slog.Attr {
	switch attr.Value.Kind() {
	case slog.KindGroup:
		h.wrapAttrs(attr.Value.Group())
		return attr
	case slog.KindAny, slog.KindLogValuer:
		switch msg := attr.Value.Any().(type) {
		case proto.Message:
			attr.Value = slog.AnyValue(Valuer{
				Message: msg,
				options: h.options,
			})
		case Valuer:
			attr.Value = slog.AnyValue(Valuer{
				Message: msg.Message,
				options: mergeOptions(h.options, msg.options),
			})
		}
		return attr
	default:
		return attr
	}
}
