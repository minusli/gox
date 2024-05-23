package xsync

import (
	"github.com/minusli/gox/xslice"
)

type ChunkExecutor[Req any, Resp any] struct {
	chunks [][]Req
}

func (ce *ChunkExecutor[Req, Resp]) Chunk(reqs []Req, size int) *ChunkExecutor[Req, Resp] {
	ce.chunks = xslice.Chunk(reqs, size)
	return ce
}

func (ce *ChunkExecutor[Req, Resp]) Execute(task func(chunk []Req) ([]Resp, error)) ([]Resp, error) {
	var resps []Resp

	for _, chunk := range ce.chunks {
		_resps, err := task(chunk)
		if err != nil {
			return nil, err
		}
		resps = append(resps, _resps...)
	}

	return resps, nil
}

func (ce *ChunkExecutor[Req, Resp]) AsyncExecute(task func(chunk []Req) ([]Resp, error)) ([]Resp, error) {
	resps := Slice[Resp]{}

	wg := WaitGroup{}
	for _, chunk := range ce.chunks {
		_chunk := chunk

		wg.Go(func() error {
			_resp, err := task(_chunk)
			if err != nil {
				return err
			}
			resps.Append(_resp...)
			return nil
		})
	}
	if err := wg.Wait(); err != nil {
		return nil, err
	}

	return resps.ToSlice(), nil
}
