package xpack

import (
	"context"
	"testing"
)

type Count struct {
	count int
}

type Result struct {
	result int
}

type Add struct {
}

func (a *Add) Load(_ context.Context, data *Count) error {
	data.count += 1
	return nil
}
func (a *Add) IgnoreError() bool {
	return true
}

type Divide struct{}

func (d *Divide) Assemble(_ context.Context, data *Count) (*Result, error) {
	return &Result{result: data.count / 10}, nil
}

func TestPack(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		p := new(Pack[Count, Result])
		for i := 0; i < 1000000; i++ {
			p.Load(i, &Add{})
		}
		p.Assemble(&Divide{})
		result, err := p.Do(context.Background(), &Count{})
		if err != nil {
			t.Errorf("pack.do error: err=%v", err)
			return
		}
		if result.result != 100000 {
			t.Errorf("pack.do failed: got=%v, want=%v", result.result, 10000)
		}
	})
}
