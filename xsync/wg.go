package xsync

import (
	"errors"
	"fmt"
	"sync"
)

type WaitGroup struct {
	parallel chan bool
	wg       sync.WaitGroup
	mutex    sync.Mutex
	err      error
}

func (wg *WaitGroup) WithParallel(max int) *WaitGroup {
	if max <= 0 {
		wg.parallel = nil
		return wg
	}

	wg.parallel = make(chan bool, max)
	return wg
}

func (wg *WaitGroup) Go(task func() error) *WaitGroup {
	wg.wg.Add(1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				wg.setErr(errors.New(fmt.Sprintf("panic: recover=%v", p)))
			}

			if wg.parallel != nil {
				<-wg.parallel
			}

			wg.wg.Done()
		}()

		if wg.parallel != nil {
			wg.parallel <- true
		}

		if err := task(); err != nil {
			wg.setErr(err)
		}
	}()
	return wg
}

func (wg *WaitGroup) Wait() error {
	wg.wg.Wait()
	return wg.err
}

func (wg *WaitGroup) setErr(err error) {
	if err == nil {
		return
	}

	wg.mutex.Lock()
	defer wg.mutex.Unlock()

	wg.err = err
}
