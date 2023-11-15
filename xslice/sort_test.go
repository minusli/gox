package xslice

import (
	"reflect"
	"testing"

	"github.com/minusli/gox/xtype"
)

func TestSort(t *testing.T) {
	type args[T any] struct {
		items   []T
		less    func(T, T) bool
		options []SortOption
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		{
			name: "sortByLengthDesc",
			args: args[string]{
				items:   []string{"a", "aa", "aaa"},
				less:    func(item1 string, item2 string) bool { return len(item1) < len(item2) },
				options: []SortOption{DESCSortOption},
			},
			want: []string{"aaa", "aa", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.items, tt.args.less, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortNumber(t *testing.T) {
	type args[T xtype.Number] struct {
		items   []T
		options []SortOption
	}
	type testCase[T xtype.Number] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[float64]{
		{
			name: "desc",
			args: args[float64]{
				items:   []float64{1.1, 1.2, 3.1},
				options: []SortOption{DESCSortOption},
			},
			want: []float64{3.1, 1.2, 1.1},
		},
		{
			name: "asc",
			args: args[float64]{
				items:   []float64{1.1, 1.2, 3.1},
				options: []SortOption{ASCSortOption},
			},
			want: []float64{1.1, 1.2, 3.1},
		},
		{
			name: "nil",
			args: args[float64]{
				items:   nil,
				options: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortNumber(tt.args.items, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortString(t *testing.T) {
	type args struct {
		items   []string
		options []SortOption
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "desc",
			args: args{
				items:   []string{"1.1", "1.2", "3.1"},
				options: []SortOption{DESCSortOption},
			},
			want: []string{"3.1", "1.2", "1.1"},
		},
		{
			name: "asc",
			args: args{
				items:   []string{"1.1", "1.2", "3.1"},
				options: []SortOption{ASCSortOption},
			},
			want: []string{"1.1", "1.2", "3.1"},
		},
		{
			name: "nil",
			args: args{
				items:   nil,
				options: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortString(tt.args.items, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortString() = %v, want %v", got, tt.want)
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
