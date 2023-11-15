package xslice

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/minusli/gox/xptr"
)

func TestDedup(t *testing.T) {
	type args[T any] struct {
		items []T
		opts  []Option[DedupOptions[T]]
	}
	type testCase[T any] struct {
		name string
		want []T
		args args[T]
	}
	tests := []testCase[string]{
		{
			name: "string-ok",
			args: args[string]{
				items: []string{"a", "a", "b", "c", "a", "b", "c"},
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "string-nil",
			args: args[string]{
				items: nil,
			},
			want: nil,
		},
		{
			name: "string-blank",
			args: args[string]{
				items: []string{},
			},
			want: []string{},
		},
		{
			name: "string-with-keyFn",
			args: args[string]{
				items: []string{"a", "a", "b", "c", "a", "b", "c"},
				opts: []Option[DedupOptions[string]]{WithDedupKeyFn(func(s string) string {
					return s
				})},
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "string-with-eqFn",
			args: args[string]{
				items: []string{"a", "a", "b", "c", "a", "b", "c"},
				opts: []Option[DedupOptions[string]]{WithDedupEqFn(func(item1, item2 string) bool {
					return item1 == item2
				})},
			},
			want: []string{"a", "b", "c"},
		},
	}

	tests1 := []testCase[int]{
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
			name: "int-blank",
			args: args[int]{
				items: []int{},
			},
			want: []int{},
		},
		{
			name: "int-with-keyFn",
			args: args[int]{
				items: []int{1, 1, 2, 3, 1, 2, 3},
				opts: []Option[DedupOptions[int]]{WithDedupKeyFn(func(s int) string {
					return strconv.FormatInt(int64(s), 10)
				})},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "int-with-eqFn",
			args: args[int]{
				items: []int{1, 1, 2, 3, 1, 2, 3},
				opts: []Option[DedupOptions[int]]{WithDedupEqFn(func(item1, item2 int) bool {
					return item1 == item2
				})},
			},
			want: []int{1, 2, 3},
		},
	}

	type person struct {
		name string
	}

	tests2 := []testCase[person]{
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
			name: "struct-blank",
			args: args[person]{
				items: []person{},
			},
			want: []person{},
		},
		{
			name: "struct-with-keyFn",
			args: args[person]{
				items: []person{{"1"}, {"1"}, {"2"}, {"3"}, {"1"}, {"2"}, {"3"}},
				opts: []Option[DedupOptions[person]]{WithDedupKeyFn(func(s person) string {
					return s.name
				})},
			},
			want: []person{{"1"}, {"2"}, {"3"}},
		},
		{
			name: "struct-with-eqFn",
			args: args[person]{
				items: []person{{"1"}, {"1"}, {"2"}, {"3"}, {"1"}, {"2"}, {"3"}},
				opts: []Option[DedupOptions[person]]{WithDedupEqFn(func(item1, item2 person) bool {
					return item1.name == item2.name
				})},
			},
			want: []person{{"1"}, {"2"}, {"3"}},
		},
	}

	ptr1 := xptr.ToPtr(1)
	ptr2 := xptr.ToPtr(2)
	ptr3 := xptr.ToPtr(3)

	tests3 := []testCase[*int]{
		{
			name: "ptr-int-ok",
			args: args[*int]{
				items: []*int{ptr1, ptr1, ptr2, ptr3, ptr1, ptr2, ptr3},
			},
			want: []*int{ptr1, ptr2, ptr3},
		},
		{
			name: "ptr-int-nil",
			args: args[*int]{
				items: nil,
			},
			want: nil,
		},
		{
			name: "ptr-int-blank",
			args: args[*int]{
				items: []*int{},
			},
			want: []*int{},
		},
		{
			name: "ptr-int-with-keyFn",
			args: args[*int]{
				items: []*int{ptr1, ptr1, ptr2, ptr3, ptr1, ptr2, ptr3},
				opts: []Option[DedupOptions[*int]]{WithDedupKeyFn(func(s *int) string {
					return strconv.FormatInt(int64(*s), 10)
				})},
			},
			want: []*int{ptr1, ptr2, ptr3},
		},
		{
			name: "ptr-int-with-eqFn",
			args: args[*int]{
				items: []*int{ptr1, ptr1, ptr2, ptr3, ptr1, ptr2, ptr3},
				opts: []Option[DedupOptions[*int]]{WithDedupEqFn(func(item1, item2 *int) bool {
					return *item1 == *item2
				})},
			},
			want: []*int{ptr1, ptr2, ptr3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dedup(tt.args.items, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dedup() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dedup(tt.args.items, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dedup() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dedup(tt.args.items, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dedup() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests3 {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dedup(tt.args.items, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dedup() = %v, want %v", got, tt.want)
			}
		})
	}
}
