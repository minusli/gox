package xpack

import (
	"context"

	"github.com/minusli/gox/xsync"
)

type Pack[Data any, Result any] struct {
	loaders  [][]Loader[Data]
	assemble Assemble[Data, Result]
}

func (a *Pack[Data, Result]) Load(level int, loaders ...Loader[Data]) *Pack[Data, Result] {
	for level >= len(a.loaders) {
		a.loaders = append(a.loaders, []Loader[Data]{})
	}

	a.loaders[level] = append(a.loaders[level], loaders...)
	return a
}

func (a *Pack[Data, Result]) Assemble(assemble Assemble[Data, Result]) *Pack[Data, Result] {
	a.assemble = assemble
	return a
}

func (a *Pack[Data, Result]) Do(ctx context.Context, data *Data) (*Result, error) {
	for _, loaders := range a.loaders {
		wg := &xsync.WaitGroup{}
		for _, loader := range loaders {
			_loader := loader
			wg.Go(func() error { return _loader.Load(ctx, data) })
		}
		if err := wg.Wait(); err != nil {
			return nil, err
		}
	}

	result, err := a.assemble.Assemble(ctx, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}
