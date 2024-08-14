package xget

import (
	"context"

	"github.com/minusli/gox/xslice"
	"github.com/minusli/gox/xsync"
)

type MGetOption struct {
	Chunk    int
	Parallel bool
}

func WithChunk(chunk int) func(*MGetOption) {
	return func(option *MGetOption) {
		option.Chunk = chunk
	}

}

func WithParallel() func(*MGetOption) {
	return func(option *MGetOption) {
		option.Parallel = true
	}
}

func MGet[ID comparable, T any](ctx context.Context, ids []ID, mget func(context.Context, []ID) (map[ID]*T, error), options ...func(*MGetOption)) (map[ID]*T, error) {
	option := &MGetOption{}
	for _, opt := range options {
		opt(option)
	}

	if option.Chunk <= 0 { // 不分片
		return mget(ctx, ids)
	}

	if option.Parallel { // 分片&&并发
		result := xsync.Map[ID, *T]{}
		wg := xsync.WaitGroup{}
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
			return nil, err
		}

		return result.ToMap(), nil
	}

	// 分片&&串行
	result := map[ID]*T{}
	for _, chunk := range xslice.Chunk(ids, option.Chunk) {
		r, err := mget(ctx, chunk)
		if err != nil {
			return nil, err
		}

		for k, v := range r {
			result[k] = v
		}
	}

	return result, nil
}
