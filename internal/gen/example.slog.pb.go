// Code generated by protoc-gen-slog. DO NOT EDIT.

package gen

import (
	protoslog "github.com/rodaine/protoslog"
	slog "log/slog"
)

// LogValue satisfies the [slog.LogValuer] interface.
func (msg *User) LogValue() slog.Value {
	return protoslog.MessageValue(msg)
}

// LogValue satisfies the [slog.LogValuer] interface.
func (msg *Location) LogValue() slog.Value {
	return protoslog.MessageValue(msg)
}
