package xslice

import (
	"reflect"
	"testing"

	"github.com/minusli/gox/xtype"
)

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
