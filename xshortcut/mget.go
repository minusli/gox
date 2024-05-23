package xshortcut

import (
	"github.com/minusli/gox/xslice"
	"github.com/minusli/gox/xsync"
	"github.com/minusli/gox/xtype"
)

type MGet[ID xtype.ID, T any] struct {
	chunk    int
	parallel bool
}

func (mget *MGet[ID, T]) Chunk(size int) *MGet[ID, T] {
	mget.chunk = size
	return mget
}

func (mget *MGet[ID, T]) Parallel() *MGet[ID, T] {
	mget.parallel = true
	return mget
}

func (mget *MGet[ID, T]) Do(ids []ID, get func(chunk []ID) (map[ID]T, error)) (map[ID]T, error) {
	if mget.chunk <= 0 {
		return get(ids)
	}

	if mget.parallel {
		ret := xsync.Map[ID, T]{}
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
		ret := map[ID]T{}
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
