package xsync

import (
	"sync"
)

type Map[K comparable, V any] struct {
	data sync.Map
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	var zero V
	val, exists := m.data.Load(key)
	if !exists {
		return zero, false
	}

	return val.(V), true
}

func (m *Map[K, V]) Put(key K, val V) {
	m.data.Store(key, val)
}

func (m *Map[K, V]) Puts(maps ...map[K]V) {
	for _, map_ := range maps {
		for k, v := range map_ {
			m.Put(k, v)
		}
	}
}

func (m *Map[K, V]) Delete(keys ...K) {
	for _, key := range keys {
		m.data.Delete(key)
	}
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.data.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (m *Map[K, V]) ToMap() map[K]V {
	ret := make(map[K]V)
	m.Range(func(key K, value V) bool {
		ret[key] = value
		return true
	})

	return ret
}
