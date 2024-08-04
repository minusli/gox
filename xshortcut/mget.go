package xshortcut

import (
	"github.com/minusli/gox/xslice"
	"github.com/minusli/gox/xsync"
)

type MGet[K comparable, T any] struct {
	chunk    int
	parallel bool
}

func (mget *MGet[K, T]) Chunk(size int) *MGet[K, T] {
	mget.chunk = size
	return mget
}

func (mget *MGet[K, T]) Parallel() *MGet[K, T] {
	mget.parallel = true
	return mget
}

func (mget *MGet[K, T]) Do(ids []K, get func(chunk []K) (map[K]T, error)) (map[K]T, error) {
	if mget.chunk <= 0 {
		return get(ids)
	}

	if mget.parallel {
		ret := xsync.Map[K, T]{}
		wg := xsync.WaitGroup{}
		for _, chunk := range xslice.Chunk(ids, mget.chunk) {
			_chunk := chunk
			wg.Go(func() error {
				kv, err := get(_chunk)
				if err != nil {
					return err
				}

				for k, v := range kv {
					ret.Put(k, v)
				}

				return nil
			})
		}
		if err := wg.Wait(); err != nil {
			return nil, err
		}

		return ret.ToMap(), nil

	} else {
		ret := map[K]T{}
		for _, chunk := range xslice.Chunk(ids, mget.chunk) {
			kv, err := get(chunk)
			if err != nil {
				return nil, err
			}

			for k, v := range kv {
				ret[k] = v
			}
		}

		return ret, nil
	}
}
