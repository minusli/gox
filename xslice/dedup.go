package xslice

import (
	"fmt"
	"reflect"

	"github.com/minusli/gox/xtype"
)

func DedupByKey[T any](items []T, key func(item T) string) []T {
	if len(items) == 0 {
		return items
	}

	var ret []T
	m := make(map[string]bool)
	for _, item := range items {
		if _, ok := m[key(item)]; !ok {
			ret = append(ret, item)
			m[key(item)] = true
		}
	}

	return ret
}

func DedupByEq[T any](items []T, eq func(item1 T, item2 T) bool) []T {
	if len(items) == 0 {
		return items
	}

	var ret []T
	for _, target := range items {
		exists := false
		for _, item := range ret {
			if eq(target, item) {
				exists = true
				break
			}
		}
		if !exists {
			ret = append(ret, target)
		}
	}

	return ret
}

func DedupString(items []string) []string {
	return DedupByKey(items, func(item string) string {
		return item
	})
}

func DedupNumber[T xtype.Number](items []T) []T {
	return DedupByKey(items, func(item T) string {
		return fmt.Sprintf("%v", item)
	})
}

func DedupPtr[T any](items []*T) []*T {
	return DedupByKey(items, func(item *T) string {
		return fmt.Sprintf("%v", item)
	})
}

func DedupWithDeepEqual[T any](items []T) []T {
	return DedupByEq(items, func(item1 T, item2 T) bool {
		return reflect.DeepEqual(item1, item2)
	})
}
