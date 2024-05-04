package xmath

import (
	"reflect"
	"testing"
)

func TestMinMax(t *testing.T) {
	t.Run("Min()#1", func(t *testing.T) {
		if got := Min(3, 2, 1); !reflect.DeepEqual(got, 1) {
			t.Errorf("Min error: got = %v, want = 1", got)
		}
	})

	t.Run("Max()#1", func(t *testing.T) {
		if got := Max(1, 2, 3); !reflect.DeepEqual(got, 3) {
			t.Errorf("Max error: got = %v, want = 3", got)
		}
	})
}
