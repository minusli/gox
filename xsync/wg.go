package xsync

import (
	"sync"
)

type WaitGroup struct {
	Parallel chan bool
	wg       sync.WaitGroup
	err      error
}

func (wg *WaitGroup) Go(task func() error) *WaitGroup {
	if wg.Parallel != nil {
		wg.Parallel <- true
	}
	wg.wg.Add(1)
	go func() {
		defer func() {
			wg.wg.Done()
			if wg.Parallel != nil {
				<-wg.Parallel
			}
		}()

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
