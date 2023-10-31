package xslice

import (
	"github.com/minusli/gox/xtype"
)

func Filter[T any](items []T, filter func(T) bool) []T {
	var retItems []T
	for _, item := range items {
		if filter(item) {
			retItems = append(retItems, item)
		}
	}

	return retItems
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
