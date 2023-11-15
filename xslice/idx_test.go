package xslice

import (
	"reflect"
	"testing"

	"github.com/minusli/gox/xptr"
)

func TestFirst(t *testing.T) {
	type args[T any] struct {
		items []T
	}
	type testCase[T any] struct {
		name      string
		args      args[T]
		wantRet   T
		wantPanic bool
	}
	testsForStrPtr := []testCase[*string]{
		{
			name: "StrPtr-ok",
			args: args[*string]{items: []*string{
				xptr.ToPtr("a"),
				xptr.ToPtr("b"),
				xptr.ToPtr("c"),
			}},
			wantRet: xptr.ToPtr("a"),
		},
		{
			name:      "StrPtr-blank",
			args:      args[*string]{items: []*string{}},
			wantPanic: true,
		},
		{
			name:      "StrPtr-nil",
			args:      args[*string]{},
			wantPanic: true,
		},
	}

	testsForStr := []testCase[string]{
		{
			name: "Str-ok",
			args: args[string]{items: []string{
				"a",
				"b",
				"c",
			}},
			wantRet: "a",
		},
		{
			name:      "Str-blank",
			args:      args[string]{items: []string{}},
			wantPanic: true,
		},
		{
			name:      "Str-nil",
			args:      args[string]{},
			wantPanic: true,
		},
	}
	for _, tt := range testsForStrPtr {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if e := recover(); e != nil && !tt.wantPanic {
					t.Errorf("First() panic = %v, wantPanic %v", e, tt.wantPanic)
				}
			}()

			if gotRet := First(tt.args.items); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("First() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
	for _, tt := range testsForStr {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if e := recover(); e != nil && !tt.wantPanic {
					t.Errorf("First() panic = %v, wantPanic %v", e, tt.wantPanic)
				}
			}()

			if gotRet := First(tt.args.items); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("First() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestLast(t *testing.T) {
	type args[T any] struct {
		items []T
	}
	type testCase[T any] struct {
		name      string
		args      args[T]
		wantRet   T
		wantPanic bool
	}
	testsForStrPtr := []testCase[*string]{
		{
			name: "StrPtr-ok",
			args: args[*string]{items: []*string{
				xptr.ToPtr("a"),
				xptr.ToPtr("b"),
				xptr.ToPtr("c"),
			}},
			wantRet: xptr.ToPtr("c"),
		},
		{
			name:      "StrPtr-blank",
			args:      args[*string]{items: []*string{}},
			wantPanic: true,
		},
		{
			name:      "StrPtr-nil",
			args:      args[*string]{},
			wantPanic: true,
		},
	}

	testsForStr := []testCase[string]{
		{
			name: "Str-ok",
			args: args[string]{items: []string{
				"a",
				"b",
				"c",
			}},
			wantRet: "c",
		},
		{
			name:      "Str-blank",
			args:      args[string]{items: []string{}},
			wantPanic: true,
		},
		{
			name:      "Str-nil",
			args:      args[string]{},
			wantPanic: true,
		},
	}
	for _, tt := range testsForStrPtr {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if e := recover(); e != nil && !tt.wantPanic {
					t.Errorf("Last() panic = %v, wantPanic %v", e, tt.wantPanic)
				}
			}()

			if gotRet := Last(tt.args.items); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("Last() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
	for _, tt := range testsForStr {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if e := recover(); e != nil && !tt.wantPanic {
					t.Errorf("Last() panic = %v, wantPanic %v", e, tt.wantPanic)
				}
			}()

			if gotRet := Last(tt.args.items); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("Last() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestMid(t *testing.T) {
	type args[T any] struct {
		items []T
	}
	type testCase[T any] struct {
		name      string
		args      args[T]
		wantLeft  T
		wantRight T
		wantPanic bool
	}
	tests := []testCase[int]{
		{
			name:      "int-odd",
			args:      args[int]{items: []int{1, 2, 3}},
			wantLeft:  2,
			wantRight: 2,
		},
		{
			name:      "int-even",
			args:      args[int]{items: []int{1, 2, 3, 4}},
			wantLeft:  2,
			wantRight: 3,
		},
		{
			name:      "int-blank",
			args:      args[int]{items: []int{}},
			wantPanic: true,
		},
		{
			name:      "int-nil",
			args:      args[int]{},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if e := recover(); e != nil && !tt.wantPanic {
					t.Errorf("Mid() panic = %v, wantPanic %v", e, tt.wantPanic)
				}
			}()

			gotLeft, gotRight := Mid(tt.args.items)
			if !reflect.DeepEqual(gotLeft, tt.wantLeft) {
				t.Errorf("Mid() gotLeft = %v, want %v", gotLeft, tt.wantLeft)
			}
			if !reflect.DeepEqual(gotRight, tt.wantRight) {
				t.Errorf("Mid() gotRight = %v, want %v", gotRight, tt.wantRight)
			}
		})
	}
}
