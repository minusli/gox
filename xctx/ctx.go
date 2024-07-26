package xctx

import (
	"context"

	"github.com/minusli/gox/xsync"
)

//goland:noinspection GoSnakeCaseUsage
const (
	CTX_KEY_KVS = ":xctx-kvs:"
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

func WithKV(ctx context.Context, kvs map[string]interface{}) context.Context {
	value := new(xsync.Map[string, interface{}])
	value.Puts(KVs(ctx), kvs)
	return context.WithValue(ctx, CTX_KEY_KVS, value)
}

func ValueKV[T any](ctx context.Context, key string, default_ T) T {
	kvs, ok := ctx.Value(CTX_KEY_KVS).(*xsync.Map[string, interface{}])
	if !ok {
		return default_
	}

	iface, exists := kvs.Get(key)
	if !exists {
		return default_
	}

	value, ok := iface.(T)
	if !ok {
		return default_
	}

	return value
}

func KVs(ctx context.Context) map[string]interface{} {
	if kvs, ok := ctx.Value(CTX_KEY_KVS).(*xsync.Map[string, interface{}]); ok {
		return kvs.ToMap()
	}

	return nil
}
