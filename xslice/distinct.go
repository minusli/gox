package xslice

import (
	"fmt"

	"github.com/minusli/gox/xtype"
)

func Distinct[T any](items []T, eq func(a, b T) bool) []T {
	var ret []T

	for _, item := range items {
		exists := false
		for _, _item := range ret {
			if eq(_item, item) {
				exists = true
				break
			}
		}

		if !exists {
			ret = append(ret, item)
		}
	}

	return ret
}

func DistinctByKey[T any, K comparable](items []T, key func(a T) K) []T {
	var ret []T

	m := map[K]bool{}
	for _, item := range items {
		if _, exists := m[key(item)]; !exists {
			ret = append(ret, item)
			m[key(item)] = true
		}
	}

	return ret
}

func DistinctStr[T xtype.String](items []T) []T {
	return DistinctByKey(items, func(a T) T {
		return a
	})
}

func DistinctNum[T xtype.Number](items []T) []T {
	return DistinctByKey(items, func(a T) T {
		return a
	})
}

func DistinctPtr[T any](items []*T) []*T {
	return DistinctByKey(items, func(a *T) string {
		return fmt.Sprintf("%p", a)
	})
}
