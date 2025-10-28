package xslice

import (
	"reflect"
)

func DistinctAny[T any](items []T, eq func(a, b T) bool) []T {
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

func Distinct[T comparable](items []T) []T {
	var ret []T
	exists := map[T]bool{}
	for _, item := range items {
		if !exists[item] {
			ret = append(ret, item)
			exists[item] = true
		}
	}

	return ret
}

func DistinctWithDeepEqual[T any](items []T) []T {
	return DistinctAny(items, func(a, b T) bool { return reflect.DeepEqual(a, b) })
}
