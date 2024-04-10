package xctx

import (
	"context"
	"reflect"
	"testing"
)

func TestCtxGet(t *testing.T) {
	t.Run("Value()#1", func(t *testing.T) {
		ctx := With(context.Background(), "key", "hello")
		if got := Value(ctx, "key", ""); !reflect.DeepEqual(got, "hello") {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Value()#2", func(t *testing.T) {
		ctx := With(context.Background(), "key", "hello")
		if got := Value(ctx, "key_unknown", "default"); !reflect.DeepEqual(got, "default") {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Value()#3", func(t *testing.T) {
		ctx := With(context.Background(), "key", "hello")
		if got := Value(ctx, "key", 1); !reflect.DeepEqual(got, 1) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
}
