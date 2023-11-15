package xslice

import (
	"reflect"
	"testing"

	"github.com/minusli/gox/xptr"
)

func TestContains(t *testing.T) {
	type args[T any] struct {
		items  []T
		target T
		opts   []Option[ContainOptions[T]]
	}
	type testCase[T any] struct {
		name string
		want bool
		args args[T]
	}
	tests := []testCase[string]{
		{
			name: "string-ok",
			args: args[string]{
				items:  []string{"a", "a", "b", "c", "a", "b", "c"},
				target: "a",
			},
			want: true,
		},
		{
			name: "string-no",
			args: args[string]{
				items:  []string{"a", "a", "b", "c", "a", "b", "c"},
				target: "d",
			},
			want: false,
		},
		{
			name: "string-nil",
			args: args[string]{
				items:  nil,
				target: "a",
			},
			want: false,
		},
		{
			name: "string-blank",
			args: args[string]{
				items:  []string{},
				target: "a",
			},
			want: false,
		},
		{
			name: "string-with-eqFn",
			args: args[string]{
				items: []string{"a", "a", "b", "c", "a", "b", "c"},
				opts: []Option[ContainOptions[string]]{WithContainEqFn(func(item1, item2 string) bool {
					return item1 == item2
				})},
				target: "a",
			},
			want: true,
		},
	}

	tests1 := []testCase[int]{
		{
			name: "int-ok",
			args: args[int]{
				items:  []int{1, 1, 2, 3, 1, 2, 3},
				target: 1,
			},
			want: true,
		},
		{
			name: "int-no",
			args: args[int]{
				items:  []int{1, 1, 2, 3, 1, 2, 3},
				target: 4,
			},
			want: false,
		},
		{
			name: "int-nil",
			args: args[int]{
				items:  nil,
				target: 1,
			},
			want: false,
		},
		{
			name: "int-blank",
			args: args[int]{
				items:  []int{},
				target: 1,
			},
			want: false,
		},
		{
			name: "int-with-eqFn",
			args: args[int]{
				items: []int{1, 1, 2, 3, 1, 2, 3},
				opts: []Option[ContainOptions[int]]{WithContainEqFn(func(item1, item2 int) bool {
					return item1 == item2
				})},
				target: 1,
			},
			want: true,
		},
	}

	type person struct {
		name string
	}

	tests2 := []testCase[person]{
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
		{
			name: "struct-with-eqFn",
			args: args[person]{
				items: []person{{"1"}, {"1"}, {"2"}, {"3"}, {"1"}, {"2"}, {"3"}},
				opts: []Option[ContainOptions[person]]{WithContainEqFn(func(item1, item2 person) bool {
					return item1.name == item2.name
				})},
				target: person{"1"},
			},
			want: true,
		},
	}

	ptr1 := xptr.ToPtr(1)
	ptr2 := xptr.ToPtr(2)
	ptr3 := xptr.ToPtr(3)

	tests3 := []testCase[*int]{
		{
			name: "ptr-int-ok",
			args: args[*int]{
				items:  []*int{ptr1, ptr1, ptr2, ptr3, ptr1, ptr2, ptr3},
				target: ptr1,
			},
			want: true,
		},
		{
			name: "ptr-int-no",
			args: args[*int]{
				items:  []*int{ptr1, ptr1, ptr2, ptr3, ptr1, ptr2, ptr3},
				target: xptr.ToPtr(1),
			},
			want: false,
		},
		{
			name: "ptr-int-nil",
			args: args[*int]{
				items:  nil,
				target: ptr1,
			},
			want: false,
		},
		{
			name: "ptr-int-blank",
			args: args[*int]{
				items:  []*int{},
				target: ptr1,
			},
			want: false,
		},
		{
			name: "ptr-int-with-eqFn",
			args: args[*int]{
				items: []*int{ptr1, ptr1, ptr2, ptr3, ptr1, ptr2, ptr3},
				opts: []Option[ContainOptions[*int]]{WithContainEqFn(func(item1, item2 *int) bool {
					return *item1 == *item2
				})},
				target: ptr1,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.items, tt.args.target, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.items, tt.args.target, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.items, tt.args.target, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests3 {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.items, tt.args.target, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
