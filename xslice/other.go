package xslice

import (
	"github.com/minusli/gox/xtype"
)

func Merge[T any](arrays ...[]T) []T {
	var ret []T
	for _, array := range arrays {
		ret = append(ret, array...)
	}

	return ret
}

func Sum[V xtype.Number](items []V) V {
	var ret V
	for _, item := range items {
		ret += item
	}
	return ret
}
