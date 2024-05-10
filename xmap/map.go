package xmap

import (
	"github.com/minusli/gox/xtype"
)

func Values[K xtype.Key, V any](m map[K]V) []V {
	var items []V
	for _, value := range m {
		items = append(items, value)
	}

	return items
}

func Keys[K xtype.Key, V any](m map[K]V) []K {
	var items []K
	for key, _ := range m {
		items = append(items, key)
	}

	return items
}

func Invert[K xtype.Key, V xtype.Key](m map[K]V) map[V]K {
	ret := make(map[V]K)
	for k, v := range m {
		ret[v] = k
	}

	return ret
}

func FlatInvert[K xtype.Key, V xtype.Key](m map[K][]V) map[V]K {
	ret := make(map[V]K)
	for k, vs := range m {
		for _, v := range vs {
			ret[v] = k
		}
	}

	return ret
}

func Get[K xtype.Key, V any](m map[K]V, key K, default_ V) V {
	val, exists := m[key]
	if !exists {
		return default_
	}

	return val
}

func GetIface[K xtype.Key, V any](m map[K]interface{}, key K, default_ V) V {
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
