package xget

import (
	"context"

	"github.com/minusli/gox/xslice"
	"github.com/minusli/gox/xsync"
)

type MGetOption struct {
	Chunk    int
	Parallel int
}

func WithChunk(chunk int) func(*MGetOption) {
	return func(option *MGetOption) {
		option.Chunk = chunk
	}

}

func WithParallel(max int) func(*MGetOption) {
	return func(option *MGetOption) {
		option.Parallel = max
	}
}

func MGet[ID comparable, T any](ctx context.Context, ids []ID, mget func(context.Context, []ID) (map[ID]*T, error), options ...func(*MGetOption)) (map[ID]*T, error) {
	// 去重&判空
	ids = xslice.Distinct(ids)
	if len(ids) == 0 {
		return nil, nil
	}

	// 可选参数
	option := &MGetOption{}
	for _, opt := range options {
		opt(option)
	}

	// 不分片&不并发
	if option.Chunk <= 0 {
		return mget(ctx, ids)
	}

	// 分片&不并发
	if option.Parallel <= 0 {
		result := map[ID]*T{}
		for _, chunk := range xslice.Chunk(ids, option.Chunk) {
			r, err := mget(ctx, chunk)
			if err != nil {
				return result, err
			}

			for k, v := range r {
				result[k] = v
			}
		}
		return result, nil
	}

	// 分片&&并发
	result := xsync.Map[ID, *T]{}
	wg := xsync.WaitGroup{}
	wg.WithParallel(option.Parallel)
	for _, chunk := range xslice.Chunk(ids, option.Chunk) {
		_chunk := chunk
		wg.Go(func() error {
			r, err := mget(ctx, _chunk)
			if err != nil {
				return err
			}

			result.Puts(r)
			return nil
		})
	}
	if err := wg.Wait(); err != nil {
		return result.ToMap(), err
	}

	return result.ToMap(), nil
}
