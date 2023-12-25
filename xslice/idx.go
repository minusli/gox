package xslice

import (
	"errors"
)

func First[T any](items []T) T {
	if len(items) == 0 {
		panic(errors.New("index out of bounds"))
	}

	return items[0]
}

func FirstWithDefault[T any](items []T, default_ T) T {
	if len(items) == 0 {
		return default_
	}

	return First(items)
}

func Last[T any](items []T) T {
	if len(items) == 0 {
		panic(errors.New("index out of bounds"))
	}

	return items[len(items)-1]
}

func LastWithDefault[T any](items []T, default_ T) T {
	if len(items) == 0 {
		return default_
	}

	return Last(items)
}

func Mid[T any](items []T) (left T, right T) {
	if len(items) == 0 {
		panic(errors.New("index out of bounds"))
	}

	if len(items)&1 == 1 {
		return items[len(items)/2], items[len(items)/2]
	} else {
		return items[len(items)/2-1], items[len(items)/2]
	}
}

func MidWithDefault[T any](items []T, default_ T) (left T, right T) {
	if len(items) == 0 {
		return default_, default_
	}

	return Mid(items)
}
