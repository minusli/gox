package xutils

import (
	"reflect"
	"testing"
)

func TestUtils(t *testing.T) {
	t.Run("IFElse()#1", func(t *testing.T) {
		if got := IFElse(true, 1, 0); !reflect.DeepEqual(got, 1) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("IFElse()#2", func(t *testing.T) {
		if got := IFElse(false, 1, 0); !reflect.DeepEqual(got, 0) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
}
