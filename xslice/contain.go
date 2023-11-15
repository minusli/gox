package xslice

import (
	"fmt"
	"reflect"
)

func Contains[T any](items []T, target T, opts ...Option[ContainOptions[T]]) bool {
	if len(items) == 0 {
		return false
	}

	opt := &ContainOptions[T]{}
	for _, optFn := range opts {
		optFn(opt)
	}

	if opt.eqFn == nil {
		switch reflect.ValueOf(target).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64, reflect.Pointer, reflect.String:
			opt.eqFn = func(t T, t2 T) bool {
				return fmt.Sprintf("%v", t) == fmt.Sprintf("%v", t2)
			}
		default:
			opt.eqFn = func(t T, t2 T) bool {
				return reflect.DeepEqual(t, t2)
			}
		}
	}

	return containsByEq(items, target, opt.eqFn)
}

func containsByEq[T any](items []T, target T, eq func(T, T) bool) bool {
	for _, item := range items {
		if eq(target, item) {
			return true
		}
	}

	return false
}
