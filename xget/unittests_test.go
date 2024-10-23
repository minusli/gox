package xget

import (
	"context"
	"reflect"
	"testing"

	"github.com/minusli/gox/xslice"
)

func TestMGet(t *testing.T) {
	ctx := context.Background()
	t.Run("不分片&不并发", func(t *testing.T) {
		var ids []int
		for i := 0; i < 100000; i++ {
			ids = append(ids, i)
		}
		want := xslice.ToMap(ids, func(a int) (int, *int) { return a, &a })

		got, err := MGet(ctx, ids, func(ctx context.Context, ids []int) (map[int]*int, error) {
			return xslice.ToMap(ids, func(a int) (int, *int) { return a, &a }), nil
		})

		if err != nil {
			t.Errorf("TestMGet#1 err = %v", err)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestMGet#1 got.len() = %v, want.len() %v", len(got), len(want))
		}
	})

	t.Run("分片&串行", func(t *testing.T) {
		var ids []int
		for i := 0; i < 100000; i++ {
			ids = append(ids, i)
		}
		want := xslice.ToMap(ids, func(a int) (int, *int) { return a, &a })

		got, err := MGet(ctx, ids, func(ctx context.Context, ids []int) (map[int]*int, error) {
			return xslice.ToMap(ids, func(a int) (int, *int) { return a, &a }), nil
		}, WithChunk(2))

		if err != nil {
			t.Errorf("TestMGet#2 err = %v", err)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestMGet#2 got.len() = %v, want.len() %v", len(got), len(want))
		}
	})

	t.Run("分片&并行", func(t *testing.T) {
		var ids []int
		for i := 0; i < 100000; i++ {
			ids = append(ids, i)
		}
		want := xslice.ToMap(ids, func(a int) (int, *int) { return a, &a })

		got, err := MGet(ctx, ids, func(ctx context.Context, ids []int) (map[int]*int, error) {
			return xslice.ToMap(ids, func(a int) (int, *int) { return a, &a }), nil
		}, WithChunk(2), WithParallel(50))

		if err != nil {
			t.Errorf("TestMGet#3 err = %v", err)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestMGet#3 got.len() = %v, want.len() %v", len(got), len(want))
		}
	})
}
