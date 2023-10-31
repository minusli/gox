package xslice

import (
	"sort"
	"strings"

	"code.byted.org/life_service/alliance_goods_goext/xtype"
)

type _SortOption struct {
	asc bool
}

type SortOption func(option *_SortOption)

func ASCSortOption(option *_SortOption) {
	option.asc = true
}

func DESCSortOption(option *_SortOption) {
	option.asc = false
}

func Sort[T any](items []T, less func(T, T) bool, options ...SortOption) []T {

	optional := &_SortOption{
		asc: true,
	}
	for _, option := range options {
		option(optional)
	}

	sort.Slice(items, func(i, j int) bool {
		if optional.asc {
			return less(items[i], items[j])
		} else {
			return !less(items[i], items[j])
		}
	})

	return items
}

func SortNumber[T xtype.Number](items []T, options ...SortOption) []T {
	return Sort(items, func(v1 T, v2 T) bool { return v1 < v2 }, options...)
}

func SortString(items []string, options ...SortOption) []string {
	return Sort(items, func(v1 string, v2 string) bool { return strings.Compare(v1, v2) < 0 }, options...)
}
