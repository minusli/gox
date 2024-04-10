package xslice

import (
	"github.com/minusli/gox/xtype"
)

func Reduce[T any, U any](items []T, reducer func(a T, result U) U, init U) U {
	for _, item := range items {
		init = reducer(item, init)
	}

	return init
}

func ReduceMap[T any, K xtype.Key, V any](items []T, kvFn func(a T) (K, V)) map[K]V {
	return Reduce(items, func(item T, result map[K]V) map[K]V {
		k, v := kvFn(item)
		result[k] = v
		return result
	}, map[K]V{})
}

func ReduceSum[T xtype.Number](items []T) T {
	return Reduce(items, func(a T, result T) T { return a + result }, 0)
}
