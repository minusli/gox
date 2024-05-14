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

func (m *Map[K, V]) Delete(key K) {
	m.data.Delete(key)
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.data.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}
