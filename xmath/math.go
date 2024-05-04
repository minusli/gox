package xmath

import (
	"github.com/minusli/gox/xtype"
)

func Min[T xtype.Number](a T, others ...T) T {
	min := a
	for _, a = range others {
		if a < min {
			min = a
		}
	}

	return min
}

func Max[T xtype.Number](a T, others ...T) T {
	max := a
	for _, a = range others {
		if a > max {
			max = a
		}
	}

	return max
}

func Clamp[T xtype.Number](a, min, max T) T {
	if a < min {
		return min
	}
	if a > max {
		return max
	}

	return a
}
