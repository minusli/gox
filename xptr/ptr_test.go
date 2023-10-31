package xptr

import (
	"reflect"
	"testing"
)

func TestToValue(t *testing.T) {
	type args[T any] struct {
		ptr *T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	helloStr := "hello"
	tests := []testCase[string]{
		{
			name: "string=hello",
			args: args[string]{
				ptr: &helloStr,
			},
			want: helloStr,
		},
		{
			name: "string=nil",
			args: args[string]{
				ptr: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToValue(tt.args.ptr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToPtr(t *testing.T) {
	type args[T any] struct {
		value T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want *T
	}
	helloStr := "hello"
	tests := []testCase[string]{
		{
			name: "string=hello",
			args: args[string]{
				value: helloStr,
			},
			want: &helloStr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToPtr(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}
