package xslice

import (
	"github.com/minusli/gox/xtype"
)

func Sum[V xtype.Number](items []V) V {
	var ret V
	for _, item := range items {
		ret += item
	}
	return ret
}

func Chunk[T any](items []T, size int) [][]T {
	var ret [][]T

	if size <= 0 {
		size = 1
	}

	for idx, item := range items {
		chunkIdx := idx / size

		if chunkIdx >= len(ret) {
			ret = append(ret, []T{})
		}

		ret[chunkIdx] = append(ret[chunkIdx], item)
	}

	return ret
}
