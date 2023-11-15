package xslice

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/minusli/gox/xptr"
	"github.com/minusli/gox/xtype"
)

func TestMap(t *testing.T) {
	type args[T any, U any] struct {
		items []T
		mapFn func(T) U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want []U
	}
	tests := []testCase[int, int]{
		{
			name: "+1",
			args: args[int, int]{
				items: []int{1, 2, 3, 4, 5},
				mapFn: func(item int) int {
					return item + 1
				},
			},
			want: []int{2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvTo(tt.args.items, tt.args.mapFn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapInterface(t *testing.T) {
	type args[T any] struct {
		items []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []interface{}
	}
	tests := []testCase[int]{
		{
			name: "ok",
			args: args[int]{
				items: []int{1, 2, 3},
			},
			want: []interface{}{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvToInterface(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvToInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapValue(t *testing.T) {
	type args[T any] struct {
		items []*T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "ok",
			args: args[int]{
				items: []*int{xptr.ToPtr(1), xptr.ToPtr(2), xptr.ToPtr(3), nil},
			},
			want: []int{1, 2, 3, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvToValue(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvToValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapPtr(t *testing.T) {
	type args[T any] struct {
		items []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []*T
	}
	tests := []testCase[int]{
		{
			name: "ok",
			args: args[int]{
				items: []int{1, 2, 3, 4},
			},
			want: []*int{xptr.ToPtr(1), xptr.ToPtr(2), xptr.ToPtr(3), xptr.ToPtr(4)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvToPtr(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvToPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatMap(t *testing.T) {
	type args[T any, U any] struct {
		items []T
		mapFn func(T) []U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want []U
	}
	tests := []testCase[string, string]{
		{
			name: "ok",
			args: args[string, string]{
				items: []string{"123", "456", "789"},
				mapFn: func(s string) []string {
					var ret []string
					for _, c := range s {
						ret = append(ret, string(c))
					}
					return ret
				},
			},
			want: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlatConvTo(tt.args.items, tt.args.mapFn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlatConvTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvToMap(t *testing.T) {
	type args[T any, V any, K xtype.Key] struct {
		items  []T
		kvFunc func(T) (K, V)
	}
	type testCase[T any, V any, K xtype.Key] struct {
		name string
		args args[T, V, K]
		want map[K]V
	}
	tests := []testCase[string, string, string]{
		{
			name: "k=v=self",
			args: args[string, string, string]{
				items: []string{"a", "aa", "aaa"},
				kvFunc: func(v string) (string, string) {
					return v, v
				},
			},
			want: map[string]string{
				"a":   "a",
				"aa":  "aa",
				"aaa": "aaa",
			},
		},
		{
			name: "k=v=len",
			args: args[string, string, string]{
				items: []string{"a", "aa", "aaa"},
				kvFunc: func(v string) (string, string) {
					return strconv.FormatInt(int64(len(v)), 10), v
				},
			},
			want: map[string]string{
				"1": "a",
				"2": "aa",
				"3": "aaa",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvToMap(tt.args.items, tt.args.kvFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
