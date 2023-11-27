package xslice

import (
	"testing"

	"github.com/minusli/gox/xptr"
	"github.com/minusli/gox/xtype"
)

func TestContainString(t *testing.T) {
	type args struct {
		items  []string
		target string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "string-ok",
			args: args{
				items:  []string{"a", "a", "b", "c", "a", "b", "c"},
				target: "a",
			},
			want: true,
		},
		{
			name: "string-no",
			args: args{
				items:  []string{"a", "a", "b", "c", "a", "b", "c"},
				target: "d",
			},
			want: false,
		},
		{
			name: "string-nil",
			args: args{
				items:  nil,
				target: "a",
			},
			want: false,
		},
		{
			name: "string-blank",
			args: args{
				items:  []string{},
				target: "a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainString(tt.args.items, tt.args.target); got != tt.want {
				t.Errorf("ContainString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainNumber(t *testing.T) {
	type args[T xtype.Number] struct {
		items  []T
		target T
	}
	type testCase[T xtype.Number] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "string-ok",
			args: args[int]{
				items:  []int{1, 1, 2, 3, 1, 2, 3},
				target: 1,
			},
			want: true,
		},
		{
			name: "string-no",
			args: args[int]{
				items:  []int{1, 1, 2, 3, 1, 2, 3},
				target: 4,
			},
			want: false,
		},
		{
			name: "string-nil",
			args: args[int]{
				items:  nil,
				target: 1,
			},
			want: false,
		},
		{
			name: "string-blank",
			args: args[int]{
				items:  []int{},
				target: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainNumber(tt.args.items, tt.args.target); got != tt.want {
				t.Errorf("ContainNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainPtr(t *testing.T) {
	ptr1 := xptr.ToPtr(1)
	ptr2 := xptr.ToPtr(1)
	ptr3 := xptr.ToPtr(1)

	type args[T any] struct {
		items  []*T
		target *T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "ptr-int-ok",
			args: args[int]{
				items:  []*int{ptr1, ptr1, ptr2, ptr3, ptr1, ptr2, ptr3},
				target: ptr1,
			},
			want: true,
		},
		{
			name: "ptr-int-no",
			args: args[int]{
				items:  []*int{ptr1, ptr1, ptr2, ptr3, ptr1, ptr2, ptr3},
				target: xptr.ToPtr(1),
			},
			want: false,
		},
		{
			name: "ptr-int-nil",
			args: args[int]{
				items:  nil,
				target: ptr1,
			},
			want: false,
		},
		{
			name: "ptr-int-blank",
			args: args[int]{
				items:  []*int{},
				target: ptr1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainPtr(tt.args.items, tt.args.target); got != tt.want {
				t.Errorf("ContainPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainWithDeepEqual(t *testing.T) {
	type person struct {
		name string
	}

	type args[T any] struct {
		items  []T
		target T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[person]{
		{
			name: "struct-ok",
			args: args[person]{
				items:  []person{{"1"}, {"1"}, {"2"}, {"3"}, {"1"}, {"2"}, {"3"}},
				target: person{"1"},
			},
			want: true,
		},
		{
			name: "struct-no",
			args: args[person]{
				items:  []person{{"1"}, {"1"}, {"2"}, {"3"}, {"1"}, {"2"}, {"3"}},
				target: person{"4"},
			},
			want: false,
		},
		{
			name: "struct-nil",
			args: args[person]{
				items:  nil,
				target: person{"1"},
			},
			want: false,
		},
		{
			name: "struct-blank",
			args: args[person]{
				items:  []person{},
				target: person{"1"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainWithDeepEqual(tt.args.items, tt.args.target); got != tt.want {
				t.Errorf("ContainWithDeepEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
