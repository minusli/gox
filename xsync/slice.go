package xsync

import (
	"sync"
)

type Slice[T any] struct {
	data  []T
	mutex sync.Mutex
}

func (s *Slice[T]) Append(items ...T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data = append(s.data, items...)
}

func (s *Slice[T]) ToSlice() []T {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return append([]T{}, s.data...)
}
