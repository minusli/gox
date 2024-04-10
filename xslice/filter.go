package xslice

import (
	"github.com/minusli/gox/xtype"
)

func Filter[T any](items []T, filter func(a T) bool) []T {
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

func FilterBlank[T xtype.String](items []T) []T {
	return Filter(items, func(a T) bool {
		return len(a) != 0
	})
}

func FilterZero[T xtype.Number](items []T) []T {
	return Filter(items, func(a T) bool {
		return a != 0
	})
}

func FilterNil[T any](items []*T) []*T {
	return Filter(items, func(a *T) bool {
		return a != nil
	})
}
