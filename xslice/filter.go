package xslice

import (
	"github.com/minusli/gox/xtype"
)

func Filter[T any](items []T, filter func(T) bool) []T {
	if len(items) == 0 {
		return items
	}

	var ret []T
	for _, item := range items {
		if filter(item) {
			ret = append(ret, item)
		}
	}

	return ret
}

func FilterZero[T xtype.Number](items []T) []T {
	return Filter(items, func(item T) bool {
		return item != 0
	})
}

func FilterNil[T any](items []*T) []*T {
	return Filter(items, func(item *T) bool {
		return item != nil
	})
}

func FilterBlank[T string | []byte](items []T) []T {
	return Filter(items, func(s T) bool {
		return len(s) == 0
	})
}
