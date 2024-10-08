package xmap

import (
	"fmt"

	"github.com/minusli/gox/xptr"
	"github.com/minusli/gox/xtype"
)

func Values[K comparable, V any](m map[K]V) []V {
	var items []V
	for _, value := range m {
		items = append(items, value)
	}

	return items
}

func Keys[K comparable, V any](m map[K]V) []K {
	var items []K
	for key, _ := range m {
		items = append(items, key)
	}

	return items
}

func Invert[K comparable, V comparable](m map[K]V) map[V]K {
	ret := make(map[V]K)
	for k, v := range m {
		ret[v] = k
	}

	return ret
}

func FlatInvert[K comparable, V comparable](m map[K][]V) map[V]K {
	ret := make(map[V]K)
	for k, vs := range m {
		for _, v := range vs {
			ret[v] = k
		}
	}

	return ret
}

func Get[K comparable, V any](m map[K]V, key K, default_ V) V {
	val, exists := m[key]
	if !exists {
		return default_
	}

	return val
}

func GetIface[K comparable, V any](m map[K]interface{}, key K, default_ V) V {
	iface, exists := m[key]
	if !exists {
		return default_
	}

	val, ok := iface.(V)
	if !ok {
		return default_
	}

	return val
}

func Map[K comparable, V any, VN any](m map[K]V, f func(V) VN) map[K]VN {
	ret := make(map[K]VN, len(m))
	for k, v := range m {
		ret[k] = f(v)
	}

	return ret
}

func MapStr[K comparable, V1 any](m map[K]V1) map[K]string {
	return Map(m, func(a V1) string {
		return fmt.Sprintf("%v", a)
	})
}

func MapVal[K comparable, V1 any](m map[K]*V1) map[K]V1 {
	return Map(m, func(a *V1) V1 {
		return xptr.Val(a)
	})
}

func MapPtr[K comparable, V1 any](m map[K]V1) map[K]*V1 {
	return Map(m, func(a V1) *V1 {
		return xptr.Ptr(a)
	})
}

func MapIface[K comparable, V1 any](m map[K]V1) map[K]interface{} {
	return Map(m, func(a V1) interface{} {
		return a
	})
}

func Filter[K comparable, V any](m map[K]V, f func(V) bool) map[K]V {
	ret := make(map[K]V)
	for k, v := range m {
		if !f(v) {
			continue
		}

		ret[k] = v
	}

	return ret
}

func FilterNil[K comparable, V any](m map[K]*V) map[K]*V {
	ret := make(map[K]*V)
	for k, v := range m {
		if v == nil {
			continue
		}

		ret[k] = v
	}

	return ret
}

func FilterBlank[K comparable](m map[K]string) map[K]string {
	ret := make(map[K]string)
	for k, v := range m {
		if v == "" {
			continue
		}

		ret[k] = v
	}

	return ret
}

func FilterZero[K comparable, V xtype.Number](m map[K]V) map[K]V {
	ret := make(map[K]V)
	for k, v := range m {
		if v == 0 {
			continue
		}

		ret[k] = v
	}

	return ret
}

func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	ret := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			ret[k] = v
		}
	}

	return ret
}

func Exists[K comparable, V any](m map[K]V, key K) bool {
	_, exists := m[key]
	return exists
}

func MapKV[K comparable, V any, K1 comparable, V1 any](m map[K]V, f func(K, V) (K1, V1)) map[K1]V1 {
	ret := make(map[K1]V1)
	for k, v := range m {
		k1, v1 := f(k, v)
		ret[k1] = v1
	}

	return ret
}
