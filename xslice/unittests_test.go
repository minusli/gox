package xslice

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minusli/gox/xptr"
)

func TestCal(t *testing.T) {
	t.Run("IntersectAny:int", func(t *testing.T) {
		if got := IntersectAny([]int{1, 1, 2, 2, 3, 3}, []int{2, 3, 3, 4, 4, 5, 5}, func(a, b int) bool { return a == b }); !reflect.DeepEqual(got, []int{2, 2, 3, 3}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Intersect:int", func(t *testing.T) {
		if got := Intersect([]int{1, 1, 2, 2, 3, 3}, []int{2, 3, 3, 4, 4, 5, 5}); !reflect.DeepEqual(got, []int{2, 2, 3, 3}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Intersect:str", func(t *testing.T) {
		if got := Intersect([]string{"1", "1", "2", "2", "3", "3"}, []string{"2", "3", "3", "4", "4", "5", "5"}); !reflect.DeepEqual(got, []string{"2", "2", "3", "3"}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Union:int", func(t *testing.T) {
		if got := Union([]int{1, 1, 2, 2, 3, 3}, []int{2, 3, 3, 4, 4, 5, 5}); !reflect.DeepEqual(got, []int{1, 1, 2, 2, 3, 3, 2, 3, 3, 4, 4, 5, 5}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("DiffAny:int", func(t *testing.T) {
		if got := DiffAny([]int{1, 1, 2, 2, 3, 3}, []int{2, 3, 3, 4, 4, 5, 5}, func(a, b int) bool { return a == b }); !reflect.DeepEqual(got, []int{1, 1}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Diff:int", func(t *testing.T) {
		if got := Diff([]int{1, 1, 2, 2, 3, 3}, []int{2, 3, 3, 4, 4, 5, 5}); !reflect.DeepEqual(got, []int{1, 1}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Diff:str", func(t *testing.T) {
		if got := Diff([]string{"1", "1", "2", "2", "3", "3"}, []string{"2", "3", "3", "4", "4", "5", "5"}); !reflect.DeepEqual(got, []string{"1", "1"}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
}

func TestContain(t *testing.T) {
	t.Run("Contains:int:found", func(t *testing.T) {
		if got := Contains([]int{1, 1, 2, 2, 3, 3}, 1); !reflect.DeepEqual(got, true) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Contains:int:not_found", func(t *testing.T) {
		if got := Contains([]int{1, 1, 2, 2, 3, 3}, 4); !reflect.DeepEqual(got, false) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Contains:str:found", func(t *testing.T) {
		if got := Contains([]string{"1", "1", "2", "2", "3", "3"}, "1"); !reflect.DeepEqual(got, true) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Contains:str:not_found", func(t *testing.T) {
		if got := Contains([]string{"1", "1", "2", "2", "3", "3"}, "4"); !reflect.DeepEqual(got, false) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Contains:ptr:found", func(t *testing.T) {
		a, b, c := xptr.Ptr(1), xptr.Ptr(2), xptr.Ptr(3)
		if got := Contains([]*int{a, b, c}, a); !reflect.DeepEqual(got, true) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Contains:ptr:not_found", func(t *testing.T) {
		a, b, c := xptr.Ptr(1), xptr.Ptr(2), xptr.Ptr(3)
		if got := Contains([]*int{a, b, c}, xptr.Ptr(1)); !reflect.DeepEqual(got, false) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("ContainsWithDeepEqual:ptr:found", func(t *testing.T) {
		a, b, c := xptr.Ptr(1), xptr.Ptr(2), xptr.Ptr(3)
		if got := ContainsWithDeepEqual([]*int{a, b, c}, xptr.Ptr(1)); !reflect.DeepEqual(got, true) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
}

func TestDistinct(t *testing.T) {
	t.Run("Distinct:str", func(t *testing.T) {
		if got := Distinct([]string{"1", "1", "2", "3", "2"}); !reflect.DeepEqual(got, []string{"1", "2", "3"}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Distinct:int", func(t *testing.T) {
		if got := Distinct([]int{1, 1, 2, 3, 2}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Distinct:ptr", func(t *testing.T) {
		a, b, c, d := xptr.Ptr(1), xptr.Ptr(2), xptr.Ptr(3), xptr.Ptr(3)
		if got := Distinct([]*int{a, a, b, c, b, d}); !reflect.DeepEqual(got, []*int{a, b, c, d}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("DistinctWithDeepEqual:ptr", func(t *testing.T) {
		a, b, c, d := xptr.Ptr(1), xptr.Ptr(2), xptr.Ptr(3), xptr.Ptr(3)
		if got := DistinctWithDeepEqual([]*int{a, a, b, c, b, d}); !reflect.DeepEqual(got, []*int{a, b, c}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("Filter()#1", func(t *testing.T) {
		if got := Filter([]int{1, 1, 2, 3, 2}, func(a int) bool { return a&1 == 1 }); !reflect.DeepEqual(got, []int{1, 1, 3}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("FilterBlank()#1", func(t *testing.T) {
		if got := FilterBlank([]string{"1", "1", "", "2", "3", "2"}); !reflect.DeepEqual(got, []string{"1", "1", "2", "3", "2"}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("FilterZero()#1", func(t *testing.T) {
		if got := FilterZero([]int{1, 1, 0, 2, 3, 2}); !reflect.DeepEqual(got, []int{1, 1, 2, 3, 2}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("FilterNil()#1", func(t *testing.T) {
		a, b, c, d := xptr.Ptr(1), xptr.Ptr(2), xptr.Ptr(3), xptr.Ptr(0)
		if got := FilterNil([]*int{a, a, nil, b, c, d}); !reflect.DeepEqual(got, []*int{a, a, b, c, d}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})

}

func TestIdx(t *testing.T) {
	t.Run("First()#1", func(t *testing.T) {
		if got := First([]int{1, 1, 2, 3, 2}); !reflect.DeepEqual(got, 1) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("First()#2", func(t *testing.T) {
		panicWrap := func() (err error) {
			defer func() {
				if p := recover(); p != nil {
					err = fmt.Errorf("%v", p)
				}
			}()
			First([]int{})
			return err
		}()

		if panicWrap == nil {
			t.Errorf("unittest error: error nil")
		}
	})
	t.Run("FirstWithDefault()#1", func(t *testing.T) {
		if got := FirstWithDefault([]int{}, 9); !reflect.DeepEqual(got, 9) {
			t.Errorf("unittest error: got = %v", got)
		}
	})

	t.Run("Last()#1", func(t *testing.T) {
		if got := Last([]int{1, 1, 2, 3, 2}); !reflect.DeepEqual(got, 2) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Last()#2", func(t *testing.T) {
		panicWrap := func() (err error) {
			defer func() {
				if p := recover(); p != nil {
					err = fmt.Errorf("%v", p)
				}
			}()
			Last([]int{})
			return err
		}()

		if panicWrap == nil {
			t.Errorf("unittest error: error nil")
		}
	})
	t.Run("LastWithDefault()#1", func(t *testing.T) {
		if got := LastWithDefault([]int{}, 9); !reflect.DeepEqual(got, 9) {
			t.Errorf("unittest error: got = %v", got)
		}
	})

	t.Run("Mid()#1", func(t *testing.T) {
		if got1, got2 := Mid([]int{1, 1, 2, 3, 2}); !reflect.DeepEqual(got1, 2) || !reflect.DeepEqual(got2, 2) {
			t.Errorf("unittest error: got1 = %v, got2 = %v", got1, got2)
		}
	})
	t.Run("Mid()#2", func(t *testing.T) {
		if got1, got2 := Mid([]int{1, 1, 3, 2}); !reflect.DeepEqual(got1, 1) || !reflect.DeepEqual(got2, 3) {
			t.Errorf("unittest error: got1 = %v, got2 = %v", got1, got2)
		}
	})
	t.Run("Mid()#3", func(t *testing.T) {
		panicWrap := func() (err error) {
			defer func() {
				if p := recover(); p != nil {
					err = fmt.Errorf("%v", p)
				}
			}()
			Mid([]int{})
			return err
		}()

		if panicWrap == nil {
			t.Errorf("unittest error: error nil")
		}
	})
	t.Run("MidWithDefault()#1", func(t *testing.T) {
		if got1, got2 := MidWithDefault([]int{}, 9); !reflect.DeepEqual(got1, 9) || !reflect.DeepEqual(got2, 9) {
			t.Errorf("unittest error: got1 = %v, got2 = %v", got1, got2)
		}
	})
}

func TestMap(t *testing.T) {
	t.Run("Map()#1", func(t *testing.T) {
		if got := Map([]int{1, 1, 2, 3, 2}, func(a int) int { return a + 1 }); !reflect.DeepEqual(got, []int{2, 2, 3, 4, 3}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("MapIface()#1", func(t *testing.T) {
		if got := MapIface([]int{1, 1, 2, 3, 2}); !reflect.DeepEqual(got, []interface{}{1, 1, 2, 3, 2}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("MapVal()#1", func(t *testing.T) {
		a, b, c, d := xptr.Ptr(1), xptr.Ptr(2), xptr.Ptr(3), xptr.Ptr(4)
		if got := MapVal([]*int{a, b, c, d, nil}); !reflect.DeepEqual(got, []int{1, 2, 3, 4, 0}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("MapPtr()#1", func(t *testing.T) {
		if got := MapPtr([]int{1, 2, 3, 4}); !reflect.DeepEqual(got, []*int{xptr.Ptr(1), xptr.Ptr(2), xptr.Ptr(3), xptr.Ptr(4)}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("MapLower()#1", func(t *testing.T) {
		if got := MapLower([]string{"Hello", "worLD"}); !reflect.DeepEqual(got, []string{"hello", "world"}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("MapUpper()#1", func(t *testing.T) {
		if got := MapUpper([]string{"Hello", "worLD"}); !reflect.DeepEqual(got, []string{"HELLO", "WORLD"}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("MapTrim()#1", func(t *testing.T) {
		if got := MapTrim([]string{",Hello", "worLD,", ",nice ,"}, ","); !reflect.DeepEqual(got, []string{"Hello", "worLD", "nice "}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("MapTrimSpace()#1", func(t *testing.T) {
		if got := MapTrimSpace([]string{" Hello", "worLD ", " nice, "}); !reflect.DeepEqual(got, []string{"Hello", "worLD", "nice,"}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Flat()#1", func(t *testing.T) {
		if got := Flat([][]int{{1, 2, 3}, {4, 5, 6}, {4, 7, 9}}); !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5, 6, 4, 7, 9}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("MapString()#1", func(t *testing.T) {
		if got := MapString([]int{1, 2, 3}); !reflect.DeepEqual(got, []string{"1", "2", "3"}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("MapString()#2", func(t *testing.T) {
		if got := MapString([]float64{1.0, 2.22, 3.333}); !reflect.DeepEqual(got, []string{"1", "2.22", "3.333"}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
}

func TestReduce(t *testing.T) {
	t.Run("ReduceAny:sum(a^2)", func(t *testing.T) {
		if got := ReduceAny([]int{1, 1, 2, 3, 2}, func(a int, result int) int { return a*a + result }, 0); !reflect.DeepEqual(got, 19) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("ToMap:int->selt", func(t *testing.T) {
		if got := ToMap([]int{1, 1, 2, 3, 2}, func(a int) (int, int) { return a, a }); !reflect.DeepEqual(got, map[int]int{1: 1, 2: 2, 3: 3}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("ToMapSelf:int", func(t *testing.T) {
		if got := ToMapSelf([]int{1, 1, 2, 3, 2}); !reflect.DeepEqual(got, map[int]int{1: 1, 2: 2, 3: 3}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Sum", func(t *testing.T) {
		if got := Sum([]int{}); !reflect.DeepEqual(got, 0) {
			t.Errorf("unittest error: got = %v", got)
		}
		if got := Sum([]int{1, 1, 2, 3, 2}); !reflect.DeepEqual(got, 9) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Join", func(t *testing.T) {
		if got := Join([]int{1, 1, 2, 3, 2}, "-"); !reflect.DeepEqual(got, "1-1-2-3-2") {
			t.Errorf("unittest 1 error: got = %v", got)
		}

		if got := Join([]interface{}{1, "hello", 2, "o", 2}, "-"); !reflect.DeepEqual(got, "1-hello-2-o-2") {
			t.Errorf("unittest 2 error: got = %v", got)
		}
	})
	t.Run("Max", func(t *testing.T) {
		if got := Max([]int{}); !reflect.DeepEqual(got, 0) {
			t.Errorf("unittest error: got = %v", got)
		}
		if got := Max([]int{1}); !reflect.DeepEqual(got, 1) {
			t.Errorf("unittest error: got = %v", got)
		}
		if got := Max([]int{1, 1, 2, 3, 2}); !reflect.DeepEqual(got, 3) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Min", func(t *testing.T) {
		if got := Min([]int{}); !reflect.DeepEqual(got, 0) {
			t.Errorf("unittest error: got = %v", got)
		}
		if got := Min([]int{1}); !reflect.DeepEqual(got, 1) {
			t.Errorf("unittest error: got = %v", got)
		}
		if got := Min([]int{1, 1, 2, 3, 2}); !reflect.DeepEqual(got, 1) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Avg", func(t *testing.T) {
		if got := Avg([]int{}); !reflect.DeepEqual(got, 0.0) {
			t.Errorf("unittest error: got = %v", got)
		}
		if got := Avg([]int{1}); !reflect.DeepEqual(got, 1.0) {
			t.Errorf("unittest error: got = %v", got)
		}
		if got := Avg([]int{1, 1, 2, 3, 2}); !reflect.DeepEqual(got, 1.8) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
}

func TestSort(t *testing.T) {
	t.Run("Sort()#1", func(t *testing.T) {
		list := []int{1, 2, 3, 3, 2, 1, 2, 1, 3}
		if Sort(list, func(a, b int) bool { return a < b }); !reflect.DeepEqual(list, []int{1, 1, 1, 2, 2, 2, 3, 3, 3}) {
			t.Errorf("unittest error: got = %v", list)
		}
	})
	t.Run("SortStr()#1", func(t *testing.T) {
		list := []string{"1", "2", "3", "3", "2", "1", "2", "1", "3"}
		if SortStr(list, false); !reflect.DeepEqual(list, []string{"1", "1", "1", "2", "2", "2", "3", "3", "3"}) {
			t.Errorf("unittest error: got = %v", list)
		}
	})
	t.Run("SortNum()#1", func(t *testing.T) {
		list := []int{1, 2, 3, 3, 2, 1, 2, 1, 3}
		if SortNum(list, true); !reflect.DeepEqual(list, []int{3, 3, 3, 2, 2, 2, 1, 1, 1}) {
			t.Errorf("unittest error: got = %v", list)
		}
	})
	t.Run("Reverse()#1", func(t *testing.T) {
		list := []int{1, 2, 3, 3, 2, 1, 2, 1, 3}
		if Reverse(list); !reflect.DeepEqual(list, []int{3, 1, 2, 1, 2, 3, 3, 2, 1}) {
			t.Errorf("unittest error: got = %v", list)
		}
	})
}

func TestSub(t *testing.T) {
	t.Run("Sub()#1", func(t *testing.T) {
		if got := Sub([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, 6, 1); !reflect.DeepEqual(got, []int{4, 5, 6}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Sub()#2", func(t *testing.T) {
		if got := Sub([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, -3, -6, -1); !reflect.DeepEqual(got, []int{7, 6, 5}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Sub()#3", func(t *testing.T) {
		if got := Sub([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 6, 3, 1); len(got) != 0 {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Sub()#4", func(t *testing.T) {
		if got := Sub([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, -6, -3, -1); len(got) != 0 {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Sub()#5", func(t *testing.T) {
		if got := Sub([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, -100, 100, 1); !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Sub()#6", func(t *testing.T) {
		if got := Sub([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 9, 0); !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Sub()#7", func(t *testing.T) {
		if got := Sub([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, -9, -1, 0); !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5, 6, 7, 8}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})

	t.Run("Chunk()#1", func(t *testing.T) {
		if got := Chunk([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3); !reflect.DeepEqual(got, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Chunk()#2", func(t *testing.T) {
		if got := Chunk([]int{1, 2, 3, 4, 5, 6, 7, 8}, -1); !reflect.DeepEqual(got, [][]int{{1, 2, 3, 4, 5, 6, 7, 8}}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
}

func TestZip(t *testing.T) {
	t.Run("Zip()#1", func(t *testing.T) {
		if got := Zip([]int{1, 2, 3}, []string{"a", "b", "c"}); !reflect.DeepEqual(got, map[int]string{1: "a", 2: "b", 3: "c"}) {
			t.Errorf("unittest error: got = %v", got)
		}

		if got := Zip([]int{1, 2}, []string{"a", "b", "c"}); !reflect.DeepEqual(got, map[int]string{}) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
}

func TestMake(t *testing.T) {
	t.Run("TestMake()#1", func(t *testing.T) {
		if got := Make(3); !reflect.DeepEqual(got, []int{0, 1, 2}) {
			t.Errorf("ugunittest error: got = %v", got)
		}
		if got := Make(1, 3); !reflect.DeepEqual(got, []int{1, 2}) {
			t.Errorf("unittest error: got = %v", got)
		}
		if got := Make(1, 3, 2); !reflect.DeepEqual(got, []int{1}) {
			t.Errorf("unittest error: got = %v", got)
		}

		if got := Make(); got != nil {
			t.Errorf("unittest error: got = %v", got)
		}
		if got := Make(0); got != nil {
			t.Errorf("unittest error: got = %v", got)
		}
		if got := Make(-1); got != nil {
			t.Errorf("unittest error: got = %v", got)
		}
		if got := Make(1, 1); got != nil {
			t.Errorf("unittest error: got = %v", got)
		}
		if got := Make(2, 1); got != nil {
			t.Errorf("unittest error: got = %v", got)
		}
		if got := Make(1, 2, 0); got != nil {
			t.Errorf("unittest error: got = %v", got)
		}
		if got := Make(1, 2, -1); got != nil {
			t.Errorf("unittest error: got = %v", got)
		}

	})
}
