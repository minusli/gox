package xslice

import (
	"reflect"
	"testing"

	"code.byted.org/life_service/alliance_goods_goext/xtype"
)

func TestContains(t *testing.T) {
	type args[T any] struct {
		items  []T
		target T
		eq     func(item T, target T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[string]{
		{
			name: "false",
			args: args[string]{
				items:  []string{"a", "b", "c"},
				target: "d",
				eq: func(item string, target string) bool {
					return item == target
				},
			},
			want: false,
		},
		{
			name: "true",
			args: args[string]{
				items:  []string{"a", "b", "c"},
				target: "a",
				eq: func(item string, target string) bool {
					return item == target
				},
			},
			want: true,
		},
		{
			name: "nil",
			args: args[string]{
				items:  nil,
				target: "a",
				eq: func(item string, target string) bool {
					return item == target
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.items, tt.args.target, tt.args.eq); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	type args[T any] struct {
		arrays [][]T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		{
			name: "ok",
			args: args[string]{
				arrays: [][]string{
					[]string{"a", "b", "c"},
					[]string{"d", "e", "f"},
					nil,
					[]string{"a", "e"},
				},
			},
			want: []string{"a", "b", "c", "d", "e", "f", "a", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.arrays...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDedup(t *testing.T) {
	type args[T any, K xtype.Key] struct {
		items  []T
		getKey func(T) K
	}
	type testCase[T any, K xtype.Key] struct {
		name string
		args args[T, K]
		want []T
	}
	tests := []testCase[string, string]{
		{
			name: "ok",
			args: args[string, string]{
				items: []string{"a", "b", "c", "b"},
				getKey: func(item string) string {
					return item
				},
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "nil",
			args: args[string, string]{
				items: nil,
				getKey: func(item string) string {
					return item
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dedup(tt.args.items, tt.args.getKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dedup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDedupNumber(t *testing.T) {
	type args[T xtype.Number] struct {
		items []T
	}
	type testCase[T xtype.Number] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "ok",
			args: args[int]{
				items: []int{1, 2, 3, 4, 1, 3, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "nil",
			args: args[int]{
				items: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DedupNumber(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DedupNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDedupString(t *testing.T) {
	type args struct {
		items []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "ok",
			args: args{
				items: []string{"a", "b", "c", "b"},
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "nil",
			args: args{
				items: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DedupString(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DedupString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args[T any] struct {
		items []T
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
				items: []int{1, 5, 2, 5, 8, 3, 4},
			},
			want: []int{4, 3, 8, 5, 2, 5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Reverse(tt.args.items)
			if !reflect.DeepEqual(tt.args.items, tt.want) {
				t.Errorf("Reverse() = %v, want %v", tt.args.items, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args[V xtype.Number] struct {
		items []V
	}
	type testCase[V xtype.Number] struct {
		name string
		args args[V]
		want V
	}
	tests := []testCase[float64]{
		{
			name: "ok",
			args: args[float64]{
				items: []float64{1.1, 2.3, 3.5},
			},
			want: 6.9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.items); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
