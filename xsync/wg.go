package xsync

import (
	"sync"
)

type WaitGroup struct {
	wg  sync.WaitGroup
	err error
}

func (wg *WaitGroup) Go(task func() error) *WaitGroup {
	wg.wg.Add(1)
	go func() {
		defer wg.wg.Done()
		if err := task(); err != nil {
			wg.err = err
		}
	}()
	return wg
}

func (wg *WaitGroup) Wait() error {
	wg.wg.Wait()
	return wg.err
}
