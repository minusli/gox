package xslice

import (
	"sort"
	"strings"

	"github.com/minusli/gox/xtype"
)

func Sort[T any](items []T, less func(a, b T) bool) {
	sort.Slice(items, func(i, j int) bool {
		return less(items[i], items[j])
	})
}

func SortStr(items []string, reverse bool) {
	if reverse {
		Sort(items, func(a string, b string) bool { return !(strings.Compare(a, b) < 0) })
		return
	}
	Sort(items, func(a string, b string) bool { return strings.Compare(a, b) < 0 })
}

func SortNum[T xtype.Number](items []T, reverse bool) {
	if reverse {
		Sort(items, func(a T, b T) bool { return a >= b })
		return
	}
	Sort(items, func(a T, b T) bool { return a < b })
}

func Reverse[T any](items []T) {
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
}
