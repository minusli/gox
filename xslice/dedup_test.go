package xslice

import (
	"reflect"
	"testing"

	"github.com/minusli/gox/xptr"
	"github.com/minusli/gox/xtype"
)

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
			name: "string-ok",
			args: args{
				items: []string{"a", "a", "b", "c", "a", "b", "c"},
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "string-nil",
			args: args{
				items: nil,
			},
			want: nil,
		},
		{
			name: "string-empty",
			args: args{
				items: []string{},
			},
			want: []string{},
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
			name: "int-ok",
			args: args[int]{
				items: []int{1, 1, 2, 3, 1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "int-nil",
			args: args[int]{
				items: nil,
			},
			want: nil,
		},
		{
			name: "int-empty",
			args: args[int]{
				items: []int{},
			},
			want: []int{},
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

func TestDedupPtr(t *testing.T) {
	type args[T any] struct {
		items []*T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []*T
	}

	ptr1 := xptr.ToPtr(1)
	ptr2 := xptr.ToPtr(2)
	ptr3 := xptr.ToPtr(3)
	ptr4 := xptr.ToPtr(1)

	tests := []testCase[int]{
		{
			name: "ptr-int-ok",
			args: args[int]{
				items: []*int{ptr1, ptr1, ptr2, ptr3, ptr1, ptr2, ptr3, ptr4},
			},
			want: []*int{ptr1, ptr2, ptr3, ptr4},
		},
		{
			name: "ptr-int-nil",
			args: args[int]{
				items: nil,
			},
			want: nil,
		},
		{
			name: "ptr-int-empty",
			args: args[int]{
				items: []*int{},
			},
			want: []*int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DedupPtr(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DedupPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDedupWithDeepEqual(t *testing.T) {

	type args[T any] struct {
		items []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}

	type person struct {
		name string
	}

	tests := []testCase[person]{
		{
			name: "struct-ok",
			args: args[person]{
				items: []person{{"1"}, {"1"}, {"2"}, {"3"}, {"1"}, {"2"}, {"3"}},
			},
			want: []person{{"1"}, {"2"}, {"3"}},
		},
		{
			name: "struct-nil",
			args: args[person]{
				items: nil,
			},
			want: nil,
		},
		{
			name: "struct-empty",
			args: args[person]{
				items: []person{},
			},
			want: []person{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DedupWithDeepEqual(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DedupWithDeepEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
