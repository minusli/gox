package xslice

import (
	"github.com/minusli/gox/xtype"
)

// Intersect 切片交集
func Intersect[T any](items []T, others []T, eq func(a, b T) bool) []T {
	var ret []T

	for _, item := range items {
		if Contains(others, item, eq) {
			ret = append(ret, item)
		}
	}

	return ret
}

// IntersectByKey 切片交集，使用 map 加速计算
func IntersectByKey[T any, K xtype.Key](items []T, others []T, keyFn func(a T) K) []T {
	var ret []T

	m := ReduceMap(others, func(a T) (K, T) { return keyFn(a), a })
	for _, item := range items {
		if _, exists := m[keyFn(item)]; exists {
			ret = append(ret, item)
		}
	}

	return ret
}

// Union 切片并集
func Union[T any](items []T, others ...[]T) []T {

	for _, other := range others {
		items = append(items, other...)
	}

	return items
}

// Diff 切片差集
func Diff[T any](items []T, others []T, eq func(a, b T) bool) []T {
	var ret []T

	for _, item := range items {
		if !Contains(others, item, eq) {
			ret = append(ret, item)
		}
	}

	return ret
}

// DiffByKey 切片交集，使用 map 加速计算
func DiffByKey[T any, K xtype.Key](items []T, others []T, keyFn func(a T) K) []T {
	var ret []T

	m := ReduceMap(others, func(a T) (K, T) { return keyFn(a), a })
	for _, item := range items {
		if _, exists := m[keyFn(item)]; !exists {
			ret = append(ret, item)
		}
	}

	return ret
}

func IntersectStr[T xtype.String](items []T, others []T) []T {
	return IntersectByKey(items, others, func(a T) T { return a })
}

func IntersectNum[T xtype.Number](items []T, others []T) []T {
	return IntersectByKey(items, others, func(a T) T { return a })
}

func DiffStr[T xtype.String](items []T, others []T) []T {
	return DiffByKey(items, others, func(a T) T { return a })
}

func DiffNum[T xtype.Number](items []T, others []T) []T {
	return DiffByKey(items, others, func(a T) T { return a })
}
