package xslice

import (
	"fmt"
	"strings"

	"github.com/minusli/gox/xtype"
)

func ReduceAny[T any, U any](items []T, reducer func(a T, result U) U, init U) U {
	for _, item := range items {
		init = reducer(item, init)
	}

	return init
}

func ToMap[T any, K comparable, V any](items []T, kvFn func(a T) (K, V)) map[K]V {
	return ReduceAny(items, func(a T, result map[K]V) map[K]V {
		k, v := kvFn(a)
		result[k] = v
		return result
	}, map[K]V{})
}

func ToMapSelf[T comparable](items []T) map[T]T {
	return ReduceAny(items, func(item T, result map[T]T) map[T]T {
		result[item] = item
		return result
	}, map[T]T{})
}

func Sum[T xtype.Number](items []T) T {
	return ReduceAny(items, func(a T, result T) T {
		return a + result
	}, 0)
}

func Join[T any](items []T, sep string) string {
	ret := ReduceAny(items, func(a T, result string) string {
		return result + fmt.Sprintf("%v%v", a, sep)
	}, "")

	return strings.Trim(ret, sep)
}

func Max[T xtype.Number](items []T) T {
	if len(items) == 0 {
		return 0
	}

	return ReduceAny(items, func(a T, result T) T {
		if a > result {
			return a
		}
		return result
	}, items[0])
}

func Min[T xtype.Number](items []T) T {
	if len(items) == 0 {
		return 0
	}

	return ReduceAny(items, func(a T, result T) T {
		if a < result {
			return a
		}
		return result
	}, items[0])
}

func Avg[T xtype.Number](items []T) float64 {
	if len(items) == 0 {
		return 0
	}

	return float64(Sum(items)) / float64(len(items))
}
