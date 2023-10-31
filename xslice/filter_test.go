package xslice

import (
	"reflect"
	"testing"

	"code.byted.org/life_service/alliance_goods_goext/xptr"
	"code.byted.org/life_service/alliance_goods_goext/xtype"
)

func TestFilter(t *testing.T) {
	type args[T any] struct {
		items  []T
		filter func(T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[*int]{
		{
			name: "<3",
			args: args[*int]{
				items: []*int{xptr.ToPtr(1), xptr.ToPtr(2), xptr.ToPtr(3), xptr.ToPtr(4), xptr.ToPtr(5)},
				filter: func(item *int) bool {
					return xptr.ToValue(item) < 3
				},
			},
			want: []*int{xptr.ToPtr(1), xptr.ToPtr(2)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.items, tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterZero(t *testing.T) {
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
				items: []int{1, 2, 3, 4, 5, 0, 1, 2, 3},
			},
			want: []int{1, 2, 3, 4, 5, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterZero(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterNil(t *testing.T) {
	type args[T any] struct {
		items []*T
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
				items: []*int{xptr.ToPtr(1), xptr.ToPtr(2), nil, nil, xptr.ToPtr(3), xptr.ToPtr(1)},
			},
			want: []*int{xptr.ToPtr(1), xptr.ToPtr(2), xptr.ToPtr(3), xptr.ToPtr(1)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterNil(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterNil() = %v, want %v", got, tt.want)
			}
		})
	}
}
