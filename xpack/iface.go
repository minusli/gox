package xpack

import (
	"context"
)

type Loader[Data any] interface {
	Load(ctx context.Context, data *Data) error
}

type Assemble[Data any, Result any] interface {
	Assemble(ctx context.Context, data *Data) (*Result, error)
}
