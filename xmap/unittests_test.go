package xmap

import (
	"reflect"
	"testing"

	"github.com/minusli/gox/xslice"
)

func TestMap(t *testing.T) {
	t.Run("Values()#1", func(t *testing.T) {
		got := Values(map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5})
		if xslice.SortNum(got, false); !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Keys()#1", func(t *testing.T) {
		got := Keys(map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5})
		if xslice.SortStr(got, false); !reflect.DeepEqual(got, []string{"1", "2", "3", "4", "5"}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Invert()#1", func(t *testing.T) {
		if got := Invert(map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}); !reflect.DeepEqual(got, map[int]string{1: "1", 2: "2", 3: "3", 4: "4", 5: "5"}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("FlatInvert()#1", func(t *testing.T) {
		if got := FlatInvert(map[string][]int{"1": {1, 2}, "2": {3, 4}, "3": {5, 6}}); !reflect.DeepEqual(got, map[int]string{1: "1", 2: "1", 3: "2", 4: "2", 5: "3", 6: "3"}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
}
