package xmath

import (
	"reflect"
	"testing"
)

func TestMath(t *testing.T) {
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

	t.Run("Clamp()#1", func(t *testing.T) {
		if got := Clamp(0, 1, 3); !reflect.DeepEqual(got, 1) {
			t.Errorf("Clamp error: got = %v, want = 1", got)
		}
	})

	t.Run("Clamp()#2", func(t *testing.T) {
		if got := Clamp(4, 1, 3); !reflect.DeepEqual(got, 3) {
			t.Errorf("Clamp error: got = %v, want = 3", got)
		}
	})

	t.Run("Clamp()#3", func(t *testing.T) {
		if got := Clamp(2, 1, 3); !reflect.DeepEqual(got, 2) {
			t.Errorf("Clamp error: got = %v, want = 2", got)
		}
	})
}
