package xmap

import (
	"reflect"
	"testing"

	"github.com/minusli/gox/xptr"
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
	t.Run("Get()#1", func(t *testing.T) {
		if got := Get(map[string]int{"k": 1}, "k", 9); !reflect.DeepEqual(got, 1) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Get()#1", func(t *testing.T) {
		if got := Get(map[string]int{"k": 1}, "k1", 9); !reflect.DeepEqual(got, 9) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("GetIface()#1", func(t *testing.T) {
		if got := GetIface(map[string]interface{}{"k": 1}, "k", 9); !reflect.DeepEqual(got, 1) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("GetIface()#2", func(t *testing.T) {
		if got := GetIface(map[string]interface{}{"k": 1}, "k1", 9); !reflect.DeepEqual(got, 9) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("GetIface()#3", func(t *testing.T) {
		if got := GetIface(map[string]interface{}{"k": 1}, "k", "v"); !reflect.DeepEqual(got, "v") {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Map", func(t *testing.T) {
		if got := Map(map[int]int{1: 1, 2: 2}, func(v int) int { return v + 1 }); !reflect.DeepEqual(got, map[int]int{1: 2, 2: 3}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("MapStr", func(t *testing.T) {
		if got := MapStr(map[int]int{1: 1, 2: 2}); !reflect.DeepEqual(got, map[int]string{1: "1", 2: "2"}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("MapVal", func(t *testing.T) {
		if got := MapVal(map[int]*int{1: xptr.Ptr(1), 2: xptr.Ptr(2)}); !reflect.DeepEqual(got, map[int]int{1: 1, 2: 2}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("MapPtr", func(t *testing.T) {
		if got := MapPtr(map[int]int{1: 1, 2: 2}); !reflect.DeepEqual(got, map[int]*int{1: xptr.Ptr(1), 2: xptr.Ptr(2)}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("MapIface", func(t *testing.T) {
		if got := MapIface(map[int]int{1: 1, 2: 2}); !reflect.DeepEqual(got, map[int]interface{}{1: 1, 2: 2}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Filter", func(t *testing.T) {
		if got := Filter(map[int]int{1: 1, 2: 2, 3: 3}, func(v int) bool { return v > 1 }); !reflect.DeepEqual(got, map[int]int{2: 2, 3: 3}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("FilterNil", func(t *testing.T) {
		if got := FilterNil(map[int]*int{1: nil, 2: xptr.Ptr(2)}); !reflect.DeepEqual(got, map[int]*int{2: xptr.Ptr(2)}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("FilterBlank", func(t *testing.T) {
		if got := FilterBlank(map[int]string{1: "", 2: "2"}); !reflect.DeepEqual(got, map[int]string{2: "2"}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("FilterZero", func(t *testing.T) {
		if got := FilterZero(map[int]int{1: 0, 2: 2}); !reflect.DeepEqual(got, map[int]int{2: 2}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Merge", func(t *testing.T) {
		if got := Merge(map[int]int{1: 1, 2: 2}, map[int]int{3: 3, 4: 4}, map[int]int{1: 5, 4: 5}); !reflect.DeepEqual(got, map[int]int{1: 5, 2: 2, 3: 3, 4: 5}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
}
