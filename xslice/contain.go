package xslice

import (
	"github.com/minusli/gox/xtype"
)

func Contains[T any](items []T, target T, eq func(a T, target T) bool) bool {

	for _, item := range items {
		if eq(item, target) {
			return true
		}
	}

	return false
}

func ContainsStr[T xtype.String](items []T, target T) bool {
	return Contains(items, target, func(a T, target T) bool {
		return a == target
	})
}

func ContainsNum[T xtype.Number](items []T, target T) bool {
	return Contains(items, target, func(a T, target T) bool {
		return a == target
	})
}

func ContainsPtr[T any](items []*T, target *T) bool {
	return Contains(items, target, func(a *T, target *T) bool {
		return a == target
	})
}
