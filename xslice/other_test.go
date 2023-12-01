package xslice

import (
	"reflect"
	"testing"

	"github.com/minusli/gox/xtype"
)

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

func TestChunk(t *testing.T) {
	type args[T any] struct {
		items []T
		size  int
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want [][]T
	}
	tests := []testCase[int]{
		{
			name: "ok",
			args: args[int]{
				items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
				size:  5,
			},
			want: [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14}},
		},
		{
			name: "blank",
			args: args[int]{
				items: []int{},
				size:  5,
			},
			want: nil,
		},
		{
			name: "nil",
			args: args[int]{
				items: nil,
				size:  5,
			},
			want: nil,
		},
		{
			name: "size<=0",
			args: args[int]{
				items: []int{1, 2, 3, 4, 5},
				size:  -1,
			},
			want: [][]int{{1}, {2}, {3}, {4}, {5}},
		},
		{
			name: "size>=len",
			args: args[int]{
				items: []int{1, 2, 3, 4, 5},
				size:  10,
			},
			want: [][]int{{1, 2, 3, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.items, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
