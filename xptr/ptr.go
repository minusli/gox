package xptr

import (
	"reflect"
)

func ToValue[T any](ptr *T) T {
	if ptr == nil {
		vt := reflect.TypeOf(ptr)
		return reflect.New(vt.Elem()).Elem().Interface().(T)
	}

	return *ptr
}

func ToPtr[T any](value T) *T {
	return &value
}
