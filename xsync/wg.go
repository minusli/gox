package xsync

import (
	"errors"
	"fmt"
	"sync"
)

type WaitGroup struct {
	parallel chan bool
	wg       sync.WaitGroup
	err      error
}

func (wg *WaitGroup) WithParallel(max int) *WaitGroup {
	wg.parallel = make(chan bool, max)
	return wg
}

func (wg *WaitGroup) Go(task func() error) *WaitGroup {
	if wg.parallel != nil {
		wg.parallel <- true
	}
	wg.wg.Add(1)
	go func() {
		defer func() {
			wg.wg.Done()
			if wg.parallel != nil {
				<-wg.parallel
			}
			if p := recover(); p != nil {
				wg.err = errors.New(fmt.Sprintf("panic: recover=%v", p))
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
