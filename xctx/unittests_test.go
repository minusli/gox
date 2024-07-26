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

func TestCtxKVs(t *testing.T) {
	t.Run("KVs()#1", func(t *testing.T) {
		ctx := context.Background()
		ctx1 := WithKV(ctx, map[string]interface{}{"ctx1": 1})
		ctx2 := WithKV(ctx, map[string]interface{}{"ctx2": "ctx2"})
		ctx11 := WithKV(ctx1, map[string]interface{}{"ctx11": 11})
		ctx21 := WithKV(ctx2, map[string]interface{}{"ctx21": "ctx21"})
		ctx111 := WithKV(ctx11, map[string]interface{}{"ctx111": 111})

		if got := KVs(ctx); len(got) != 0 {
			t.Errorf("unittest error: ctx.KVs got = %v", got)
		}
		if got := KVs(ctx1); !reflect.DeepEqual(got, map[string]interface{}{"ctx1": 1}) {
			t.Errorf("unittest error: ctx1.KVs got = %v", got)
		}
		if got := KVs(ctx11); !reflect.DeepEqual(got, map[string]interface{}{"ctx1": 1, "ctx11": 11}) {
			t.Errorf("unittest error: ctx11.KVs got = %v", got)
		}
		if got := KVs(ctx2); !reflect.DeepEqual(got, map[string]interface{}{"ctx2": "ctx2"}) {
			t.Errorf("unittest error: ctx2.KVs got = %v", got)
		}
		if got := KVs(ctx21); !reflect.DeepEqual(got, map[string]interface{}{"ctx2": "ctx2", "ctx21": "ctx21"}) {
			t.Errorf("unittest error: ctx21.KVs got = %v", got)
		}
		if got := KVs(ctx111); !reflect.DeepEqual(got, map[string]interface{}{"ctx1": 1, "ctx11": 11, "ctx111": 111}) {
			t.Errorf("unittest error: ctx111.KVs got = %v", got)
		}

		if got := ValueKV(ctx, "ctx", ""); !reflect.DeepEqual(got, "") {
			t.Errorf("unittest error: ctx.ValueKV got = %v", got)
		}
		if got := ValueKV(ctx1, "ctx", ""); !reflect.DeepEqual(got, "") {
			t.Errorf("unittest error: ctx1.ValueKV got = %v", got)
		}
		if got := ValueKV(ctx1, "ctx1", ""); !reflect.DeepEqual(got, "") {
			t.Errorf("unittest error: ctx1.ValueKV got = %v", got)
		}
		if got := ValueKV(ctx2, "ctx2", ""); !reflect.DeepEqual(got, "ctx2") {
			t.Errorf("unittest error: ctx2.ValueKV got = %v", got)
		}
		if got := ValueKV(ctx111, "ctx11", 0); !reflect.DeepEqual(got, 11) {
			t.Errorf("unittest error: ctx111.ValueKV got = %v", got)
		}
	})
}
