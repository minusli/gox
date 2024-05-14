package xsync

import (
	"errors"
	"reflect"
	"sort"
	"sync/atomic"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	t.Run("WaitGroup()#1", func(t *testing.T) {
		count := int32(0)
		wg := WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Go(func() error {
				time.Sleep(1 * time.Second)
				atomic.AddInt32(&count, 1)
				return nil
			})
		}
		_ = wg.Wait()
		if !reflect.DeepEqual(count, int32(1000)) {
			t.Errorf("unittest error: got = %v", count)
		}
	})
	t.Run("WaitGroup()#2", func(t *testing.T) {
		count := int32(0)
		wg := WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Go(func() error {
				time.Sleep(1 * time.Second)
				atomic.AddInt32(&count, 1)
				if count == 999 {
					return errors.New("999")
				}
				return nil
			})
		}
		if got := wg.Wait(); !reflect.DeepEqual(got, errors.New("999")) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
}

func TestMap(t *testing.T) {
	t.Run("Map()#1", func(t *testing.T) {
		m := Map[int, int]{}
		wg := WaitGroup{}
		total := 100000
		delTotal := 10000
		wg.Go(func() error {
			for i := 0; i < total; i++ {
				m.Put(i, i)
			}
			return nil
		})
		wg.Go(func() error {
			for i := 0; i < total; i++ {
				m.Put(i, i)
			}
			return nil
		})
		wg.Go(func() error {
			for i := 0; i < delTotal; i++ {
				m.Delete(i)
			}
			return nil
		})
		wg.Go(func() error {
			for i := 0; i < delTotal; i++ {
				m.Delete(i)
			}
			return nil
		})
		wg.Go(func() error {
			m.Range(func(key int, value int) bool {
				key += 1
				value += 1
				return true
			})
			return nil
		})
		wg.Go(func() error {
			for i := 0; i < total; i++ {
				if val, exists := m.Get(i); exists {
					val += 1
				}
			}
			return nil
		})

		_ = wg.Wait()

		for i := 0; i < delTotal; i++ {
			m.Delete(i)
		}

		got := m.ToMap()
		want := map[int]int{}
		for i := delTotal; i < total; i++ {
			want[i] = i
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Map()#1 got.len() = %v, want.len() %v", len(got), len(want))
		}
	})
}

func TestSlice(t *testing.T) {
	t.Run("Slice#1", func(t *testing.T) {
		s := Slice[int]{}

		wg := WaitGroup{}
		total := 100000
		wg.Go(func() error {
			for i := 0; i < total; i++ {
				s.Append(i)
			}
			return nil
		})
		wg.Go(func() error {
			for i := 0; i < total; i++ {
				s.Append(i)
			}
			return nil
		})
		_ = wg.Wait()

		got := s.ToSlice()
		var want []int
		for i := 0; i < total; i++ {
			want = append(want, i, i)
		}

		sort.Ints(got)
		sort.Ints(want)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Map()#1 got.len() = %v, want.len() %v", len(got), len(want))
		}

	})
}
