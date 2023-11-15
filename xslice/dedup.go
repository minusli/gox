package xslice

import (
	"fmt"
	"reflect"
)

func Dedup[T any](items []T, opts ...Option[DedupOptions[T]]) []T {
	if len(items) == 0 {
		return items
	}

	opt := &DedupOptions[T]{}
	for _, optFn := range opts {
		optFn(opt)
	}

	if opt.keyFn == nil && opt.eqFn == nil {
		switch reflect.ValueOf(items[0]).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64, reflect.Pointer, reflect.String:
			opt.keyFn = func(t T) string {
				return fmt.Sprintf("%v", t)
			}
		default:
			opt.eqFn = func(t T, t2 T) bool {
				return reflect.DeepEqual(t, t2)
			}
		}
	}

	if opt.keyFn != nil {
		return dedupByKey(items, opt.keyFn)
	}

	return dedupByEq(items, opt.eqFn)
}

func dedupByKey[T any](items []T, keyFn func(T) string) []T {
	var ret []T

	m := make(map[string]bool)
	for _, item := range items {
		if _, ok := m[keyFn(item)]; !ok {
			ret = append(ret, item)
			m[keyFn(item)] = true
		}
	}

	return ret
}

func dedupByEq[T any](items []T, eq func(T, T) bool) []T {
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
