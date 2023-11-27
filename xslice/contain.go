package xslice

import (
	"reflect"

	"github.com/minusli/gox/xtype"
)

func Contains[T any](items []T, target T, eq func(item T, target T) bool) bool {
	for _, item := range items {
		if eq(item, target) {
			return true
		}
	}

	return false
}

func ContainString(items []string, target string) bool {
	return Contains(items, target, func(item string, target string) bool {
		return item == target
	})
}

func ContainNumber[T xtype.Number](items []T, target T) bool {
	return Contains(items, target, func(item T, target T) bool {
		return item == target
	})
}

func ContainPtr[T any](items []*T, target *T) bool {
	return Contains(items, target, func(item *T, target *T) bool {
		return item == target
	})
}

func ContainWithDeepEqual[T any](items []T, target T) bool {
	return Contains(items, target, func(item T, target T) bool {
		return reflect.DeepEqual(item, target)
	})
}
