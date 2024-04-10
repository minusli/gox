package xctx

import (
	"context"
)

func With(ctx context.Context, key string, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

func Value[T any](ctx context.Context, key string, default_ T) T {
	iface := ctx.Value(key)
	if iface == nil {
		return default_
	}

	value, ok := iface.(T)
	if !ok {
		return default_
	}

	return value
}
