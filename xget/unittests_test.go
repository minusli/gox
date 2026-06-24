package xget

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/minusli/gox/xslice"
)

func makeIDs(total int) ([]int, map[int]*int) {
	var ids []int
	for i := 0; i < total; i++ {
		ids = append(ids, i)
	}

	want := xslice.ToMap(ids, func(a int) (int, *int) { return a, &a })
	return ids, want
}

func TestMGet(t *testing.T) {
	ctx := context.Background()

	t.Run("不分片&不并发", func(t *testing.T) {
		ids, want := makeIDs(1000)

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
		ids, want := makeIDs(1000)

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
		ids, want := makeIDs(1000)

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

	t.Run("分片&串行&报错时尽量返回已收集数据", func(t *testing.T) {
		ids := []int{1, 2, 3, 4}
		want := map[int]*int{}
		for _, id := range []int{1, 2, 3} {
			id := id
			want[id] = &id
		}

		got, err := MGet(ctx, ids, func(ctx context.Context, ids []int) (map[int]*int, error) {
			result := xslice.ToMap(ids, func(a int) (int, *int) { return a, &a })
			if xslice.Contains(ids, 3) {
				delete(result, 4)
				return result, errors.New("partial error")
			}

			return result, nil
		}, WithChunk(2))

		if err == nil || err.Error() != "partial error" {
			t.Errorf("TestMGet#4 err = %v", err)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestMGet#4 got = %v, want = %v", got, want)
		}
	})

	t.Run("分片&并行&报错时尽量返回已收集数据", func(t *testing.T) {
		ids := []int{1, 2, 3, 4}
		want := map[int]*int{}
		for _, id := range ids {
			id := id
			want[id] = &id
		}

		got, err := MGet(ctx, ids, func(ctx context.Context, ids []int) (map[int]*int, error) {
			result := xslice.ToMap(ids, func(a int) (int, *int) { return a, &a })
			if xslice.Contains(ids, 3) {
				return result, errors.New("partial error")
			}

			return result, nil
		}, WithChunk(2), WithParallel(2))

		if err == nil || err.Error() != "partial error" {
			t.Errorf("TestMGet#5 err = %v", err)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestMGet#5 got = %v, want = %v", got, want)
		}
	})

	t.Run("分片&并行&多分片报错时尽量返回已收集数据", func(t *testing.T) {
		ids := []int{1, 2, 3, 4, 5, 6}
		want := map[int]*int{}
		for _, id := range ids {
			id := id
			want[id] = &id
		}

		got, err := MGet(ctx, ids, func(ctx context.Context, ids []int) (map[int]*int, error) {
			result := xslice.ToMap(ids, func(a int) (int, *int) { return a, &a })
			if xslice.Contains(ids, 3) {
				return result, errors.New("partial error 1")
			}
			if xslice.Contains(ids, 5) {
				return result, errors.New("partial error 2")
			}

			return result, nil
		}, WithChunk(2), WithParallel(3))

		if err == nil {
			t.Errorf("TestMGet#6 err = nil")
			return
		}
		if err.Error() != "partial error 1" && err.Error() != "partial error 2" {
			t.Errorf("TestMGet#6 err = %v", err)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestMGet#6 got = %v, want = %v", got, want)
		}
	})
}
