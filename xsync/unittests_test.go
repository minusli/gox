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
