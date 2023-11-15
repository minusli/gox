package xslice

type Option[T any] func(t *T)

type ContainOptions[T any] struct {
	eqFn func(T, T) bool
}

func WithContainEqFn[T any](eqFn func(T, T) bool) Option[ContainOptions[T]] {
	return func(opts *ContainOptions[T]) {
		opts.eqFn = eqFn
	}
}

type DedupOptions[T any] struct {
	keyFn func(T) string
	eqFn  func(T, T) bool
}

func WithDedupKeyFn[T any](keyFn func(T) string) Option[DedupOptions[T]] {
	return func(opts *DedupOptions[T]) {
		opts.keyFn = keyFn
	}
}

func WithDedupEqFn[T any](eqFn func(T, T) bool) Option[DedupOptions[T]] {
	return func(opts *DedupOptions[T]) {
		opts.eqFn = eqFn
	}
}
