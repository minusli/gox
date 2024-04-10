package xslice

import (
	"strings"

	"github.com/minusli/gox/xptr"
)

func Map[T any, U any](items []T, convFn func(a T) U) []U {
	var retItems []U
	for _, item := range items {
		retItems = append(retItems, convFn(item))
	}

	return retItems
}

func MapIface[T any](items []T) []interface{} {
	return Map(items, func(a T) interface{} {
		return a
	})
}

func MapVal[T any](items []*T) []T {
	return Map(items, func(a *T) T {
		return xptr.Val(a)
	})
}

func MapPtr[T any](items []T) []*T {
	return Map(items, func(a T) *T {
		return xptr.Ptr(a)
	})
}

func MapLower(items []string) []string {
	return Map(items, func(a string) string {
		return strings.ToLower(a)
	})
}

func MapUpper(items []string) []string {
	return Map(items, func(a string) string {
		return strings.ToUpper(a)
	})
}

func MapTrim(items []string, cutset string) []string {
	return Map(items, func(a string) string {
		return strings.Trim(a, cutset)
	})
}

func MapTrimSpace(items []string) []string {
	return MapTrim(items, " ")
}

func Flat[T any](items [][]T) []T {
	var ret []T

	for _, _items := range items {
		ret = append(ret, _items...)
	}

	return ret
}
