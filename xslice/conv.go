package xslice

import (
	"strings"

	"github.com/minusli/gox/xptr"
	"github.com/minusli/gox/xtype"
)

func ConvTo[T any, U any](items []T, convFn func(item T) U) []U {
	var retItems []U
	for _, item := range items {
		retItems = append(retItems, convFn(item))
	}

	return retItems
}

func ConvToInterface[T any](items []T) []interface{} {
	return ConvTo(items, func(item T) interface{} {
		return item
	})
}

func ConvToValue[T any](items []*T) []T {
	return ConvTo(items, func(item *T) T {
		return xptr.ToValue(item)
	})
}

func ConvToPtr[T any](items []T) []*T {
	return ConvTo(items, func(item T) *T {
		return xptr.ToPtr(item)
	})
}

func ConvToLower(items []string) []string {
	return ConvTo(items, func(item string) string {
		return strings.ToLower(item)
	})
}

func ConvToUpper(items []string) []string {
	return ConvTo(items, func(item string) string {
		return strings.ToUpper(item)
	})
}

func FlatConvTo[T any, U any](items []T, flatFn func(T) []U) []U {
	var retItems []U
	for _, item := range items {
		retItems = append(retItems, flatFn(item)...)
	}

	return retItems
}

func ConvToMap[T any, K xtype.Key, V any](items []T, kvFn func(T) (K, V)) map[K]V {
	retMap := make(map[K]V)
	for _, item := range items {
		k, v := kvFn(item)
		retMap[k] = v
	}

	return retMap
}

func ConvByTrim(items []string, cutset string) []string {
	return ConvTo(items, func(item string) string {
		return strings.Trim(item, cutset)
	})
}

func ConvByTrimSpace(items []string) []string {
	return ConvByTrim(items, " ")
}
