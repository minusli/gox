package xslice

import (
	"code.byted.org/life_service/alliance_goods_goext/xtype"
)

func Contains[T any](items []T, target T, eq func(item T, target T) bool) bool {
	for _, item := range items {
		if eq(item, target) {
			return true
		}
	}
	return false
}

func Merge[T any](arrays ...[]T) []T {
	var ret []T
	for _, array := range arrays {
		ret = append(ret, array...)
	}

	return ret
}

func Dedup[T any, K xtype.Key](items []T, getKey func(T) K) []T {
	m := make(map[K]bool)
	var ret []T
	for _, item := range items {
		key := getKey(item)
		if _, ok := m[key]; ok {
			continue
		}

		ret = append(ret, item)
		m[key] = true
	}

	return ret
}

func DedupNumber[T xtype.Number](items []T) []T {
	return Dedup(items, func(item T) T {
		return item
	})
}

func DedupString(items []string) []string {
	return Dedup(items, func(item string) string {
		return item
	})
}

func Reverse[T any](items []T) {
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
}

func Sum[V xtype.Number](items []V) V {
	var ret V
	for _, item := range items {
		ret += item
	}
	return ret
}

func SubSlice[T any](items []T, start, end, step int) []T {
	total := len(items)
	if start < 0 {
		start = 0
	}
	if start > total {
		start = total
	}
	if end < 0 {
		end = total
	}
	if end > total {
		end = total
	}
	if step <= 0 {
		step = 1
	}
	if start >= end {
		return nil
	}

	var ret []T
	for ; start < end; start += step {
		ret = append(ret, items[start])
	}

	return ret
}
