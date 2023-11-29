package xslice

import (
	"github.com/minusli/gox/xtype"
)

func Filter[T any](items []T, filter func(item T) bool) []T {
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

func FilterBlank(items []string) []string {
	return Filter(items, func(item string) bool {
		return item != ""
	})
}
