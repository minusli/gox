package xslice

import (
	"fmt"

	"github.com/minusli/gox/xmap"
	"github.com/minusli/gox/xtype"
)

func Sum[V xtype.Number](items []V) V {
	var ret V
	for _, item := range items {
		ret += item
	}
	return ret
}

func Chunk[T any](items []T, size int) [][]T {
	var ret [][]T

	if size <= 0 {
		size = 1
	}

	for idx, item := range items {
		chunkIdx := idx / size

		if chunkIdx >= len(ret) {
			ret = append(ret, []T{})
		}

		ret[chunkIdx] = append(ret[chunkIdx], item)
	}

	return ret
}

func DiffByKey[T any](items []T, keyFn func(item T) string, otherItems ...[]T) []T {
	others := FlatConvTo(otherItems, func(items []T) []T { return items })
	otherMap := ConvToMap(others, func(item T) (string, T) {
		return keyFn(item), item
	})

	var diffs []T
	for _, item := range items {
		if _, exists := otherMap[keyFn(item)]; !exists {
			diffs = append(diffs, item)
		}
	}

	return diffs
}

func DiffStrings(items []string, otherItems ...[]string) []string {
	return DiffByKey(items, func(item string) string { return item }, otherItems...)
}

func DiffNumbers[T xtype.Number](items []T, otherItems ...[]T) []T {
	return DiffByKey(items, func(item T) string { return fmt.Sprintf("%v", item) }, otherItems...)
}

func IntersectByKey[T any](items []T, keyFn func(item T) string, otherItems ...[]T) []T {
	m := ConvToMap(items, func(item T) (string, T) { return keyFn(item), item })
	for _, others := range otherItems {
		nextm := make(map[string]T)
		for _, other := range others {
			if _, exists := m[keyFn(other)]; exists {
				nextm[keyFn(other)] = other
			}
		}
		m = nextm
	}

	return xmap.Values(m)
}

func IntersectStrings(items []string, otherItems ...[]string) []string {
	return IntersectByKey(items, func(item string) string { return item }, otherItems...)
}

func IntersectNumbers[T xtype.Number](items []T, otherItems ...[]T) []T {
	return IntersectByKey(items, func(item T) string { return fmt.Sprintf("%v", item) }, otherItems...)
}

func UnionByKey[T any](items []T, keyFn func(item T) string, otherItems ...[]T) []T {
	m := ConvToMap(items, func(item T) (string, T) { return keyFn(item), item })
	for _, other := range FlatConvTo(otherItems, func(items []T) []T { return items }) {
		if _, exists := m[keyFn(other)]; !exists {
			items = append(items, other)
		}
	}

	return items
}

func UnionStrings(items []string, otherItems ...[]string) []string {
	return UnionByKey(items, func(item string) string { return item }, otherItems...)
}

func UnionNumber[T xtype.Number](items []T, otherItems ...[]T) []T {
	return UnionByKey(items, func(item T) string { return fmt.Sprintf("%v", item) }, otherItems...)
}
