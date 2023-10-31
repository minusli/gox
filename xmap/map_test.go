package xmap

import (
	"reflect"
	"testing"

	"github.com/minusli/gox/xslice"
	"github.com/minusli/gox/xtype"
)

func TestValues(t *testing.T) {
	type args[K xtype.Key, V any] struct {
		m map[K]V
	}
	type testCase[K xtype.Key, V any] struct {
		name string
		args args[K, V]
		want []V
	}
	tests := []testCase[string, int]{
		{
			name: "ok",
			args: args[string, int]{
				m: map[string]int{
					"a": 1,
					"b": 2,
				},
			},
			want: []int{1, 2},
		},
		{
			name: "nil",
			args: args[string, int]{
				m: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Values(tt.args.m); !reflect.DeepEqual(xslice.SortNumber(got), xslice.SortNumber(tt.want)) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeys(t *testing.T) {
	type args[K xtype.Key, V any] struct {
		m map[K]V
	}
	type testCase[K xtype.Key, V any] struct {
		name string
		args args[K, V]
		want []K
	}
	tests := []testCase[string, int]{
		{
			name: "ok",
			args: args[string, int]{
				m: map[string]int{
					"a": 1,
					"b": 2,
				},
			},
			want: []string{"a", "b"},
		},
		{
			name: "nil",
			args: args[string, int]{
				m: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Keys(tt.args.m); !reflect.DeepEqual(xslice.SortString(got), xslice.SortString(tt.want)) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvert(t *testing.T) {
	type args[K xtype.Key, V xtype.Key] struct {
		m map[K]V
	}
	type testCase[K xtype.Key, V xtype.Key] struct {
		name string
		args args[K, V]
		want map[V]K
	}
	tests := []testCase[int, string]{
		{
			name: "ok",
			args: args[int, string]{
				m: map[int]string{
					1: "a",
					2: "b",
					3: "c",
				},
			},
			want: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Invert(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Invert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatInvert(t *testing.T) {
	type args[K xtype.Key, V xtype.Key] struct {
		m map[K][]V
	}
	type testCase[K xtype.Key, V xtype.Key] struct {
		name string
		args args[K, V]
		want map[V]K
	}
	tests := []testCase[int, string]{
		{
			name: "ok",
			args: args[int, string]{
				m: map[int][]string{
					1: {"a", "d", "f"},
					2: {"b", "e"},
					3: {"c"},
				},
			},
			want: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
				"d": 1,
				"e": 2,
				"f": 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlatInvert(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlatInvert() = %v, want %v", got, tt.want)
			}
		})
	}
}
