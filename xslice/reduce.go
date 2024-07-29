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

func Sum[T xtype.Number](items []T) T {
	var sum T = 0
	for _, item := range items {
		sum += item
	}

	return sum
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

func Max[T xtype.Number](items []T) T {
	if len(items) == 0 {
		return 0
	}
	if len(items) == 0 {
		return items[0]
	}

	max := items[0]
	for _, item := range items[1:] {
		if item > max {
			max = item
		}
	}

	return max
}

func Min[T xtype.Number](items []T) T {
	if len(items) == 0 {
		return 0
	}
	if len(items) == 0 {
		return items[0]
	}

	min := items[0]
	for _, item := range items[1:] {
		if item < min {
			min = item
		}
	}

	return min
}

func Avg[T xtype.Number](items []T) float64 {
	if len(items) == 0 {
		return 0
	}

	return float64(Sum(items)) / float64(len(items))
}
