package xslice

import (
	"reflect"
	"testing"
)

func TestSubSlice(t *testing.T) {
	type args[T any] struct {
		items []T
		start int
		end   int
		step  int
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
				items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				start: 2,
				end:   8,
				step:  2,
			},
			want: []int{3, 5, 7},
		},
		{
			name: "start<0",
			args: args[int]{
				items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				start: -1,
				end:   8,
				step:  2,
			},
			want: []int{1, 3, 5, 7},
		},
		{
			name: "end<0",
			args: args[int]{
				items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				start: 2,
				end:   -1,
				step:  2,
			},
			want: []int{3, 5, 7, 9},
		},
		{
			name: "start<=end",
			args: args[int]{
				items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				start: 2,
				end:   2,
				step:  2,
			},
			want: nil,
		},
		{
			name: "start>total",
			args: args[int]{
				items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				start: 100,
				end:   8,
				step:  2,
			},
			want: nil,
		},
		{
			name: "end>total",
			args: args[int]{
				items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				start: 2,
				end:   100,
				step:  2,
			},
			want: []int{3, 5, 7, 9},
		},
		{
			name: "step<=0",
			args: args[int]{
				items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				start: 2,
				end:   8,
				step:  0,
			},
			want: []int{3, 4, 5, 6, 7, 8},
		},
		{
			name: "step>total",
			args: args[int]{
				items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				start: 2,
				end:   8,
				step:  100,
			},
			want: []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubSlice(tt.args.items, tt.args.start, tt.args.end, tt.args.step); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
