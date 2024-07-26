package xslice

import (
	"fmt"

	"github.com/minusli/gox/xtype"
)

func Reduce[T any, U any](items []T, reducer func(a T, result U) U, init U) U {
	for _, item := range items {
		init = reducer(item, init)
	}

	return init
}

func ReduceMap[T any, K comparable, V any](items []T, kvFn func(a T) (K, V)) map[K]V {
	return Reduce(items, func(item T, result map[K]V) map[K]V {
		k, v := kvFn(item)
		result[k] = v
		return result
	}, map[K]V{})
}

func ReduceMapSelf[T comparable](items []T) map[T]T {
	return Reduce(items, func(item T, result map[T]T) map[T]T {
		result[item] = item
		return result
	}, map[T]T{})
}

func ReduceSum[T xtype.Number](items []T) T {
	return Reduce(items, func(a T, result T) T { return a + result }, 0)
}

func Join[T any](items []T, sep string) string {
	if len(items) == 0 {
		return ""
	}
	if len(items) == 1 {
		return fmt.Sprintf("%v", items[0])
	}

	ret := fmt.Sprintf("%v", items[0])
	for _, item := range items[1:] {
		ret += sep + fmt.Sprintf("%v", item)
	}
	
	return ret
}
