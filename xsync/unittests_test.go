package xsync

import (
	"errors"
	"reflect"
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
		wg.Go(func() error {
			for i := 0; i < 100000; i++ {
				m.Put(i, i)
			}
			return nil
		})
		wg.Go(func() error {
			for i := 0; i < 100000; i++ {
				m.Put(i, i)
			}
			return nil
		})
		wg.Go(func() error {
			for i := 0; i < 10000; i++ {
				m.Delete(i)
			}
			return nil
		})
		wg.Go(func() error {
			for i := 0; i < 10000; i++ {
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
			for i := 0; i < 100000; i++ {
				if val, exists := m.Get(i); exists {
					val += 1
				}
			}
			return nil
		})

		_ = wg.Wait()

		wg.Go(func() error {
			for i := 0; i < 10000; i++ {
				m.Delete(i)
			}
			return nil
		})

		want := map[int]int{}
		got := map[int]int{}
		for i := 10000; i < 100000; i++ {
			want[i] = i
		}
		m.Range(func(key int, value int) bool {
			got[key] = value
			return true
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Map()#1 got = %v, want %v", got, want)
		}
	})
}
