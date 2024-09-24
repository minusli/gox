package xslice

import (
	"reflect"
)

func ContainsAny[T any](items []T, target T, eq func(a T, target T) bool) bool {

	for _, item := range items {
		if eq(item, target) {
			return true
		}
	}

	return false
}

func Contains[T comparable](items []T, target T) bool {
	return ContainsAny(items, target, func(a T, target T) bool {
		return a == target
	})
}

func ContainsWithDeepEqual[T any](items []T, target T) bool {
	return ContainsAny(items, target, func(a T, target T) bool {
		return reflect.DeepEqual(a, target)
	})
}
