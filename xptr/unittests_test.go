package xptr

import (
	"reflect"
	"testing"
)

func TestPtr(t *testing.T) {
	var nilInt *int
	var nilStr *string
	int1 := 1
	str1 := "1"
	t.Run("Ptr()#1", func(t *testing.T) {
		if got := Ptr(int1); !reflect.DeepEqual(got, &int1) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Ptr()#2", func(t *testing.T) {
		if got := Ptr(str1); !reflect.DeepEqual(got, &str1) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Val()#1", func(t *testing.T) {
		if got := Val(&int1); !reflect.DeepEqual(got, int1) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Val()#2", func(t *testing.T) {
		if got := Val(&str1); !reflect.DeepEqual(got, str1) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Val()#3", func(t *testing.T) {
		if got := Val(nilInt); !reflect.DeepEqual(got, 0) {
			t.Errorf("unittest error: got = %v", got)
		}
	})
	t.Run("Val()#4", func(t *testing.T) {
		if got := Val(nilStr); !reflect.DeepEqual(got, "") {
			t.Errorf("unittest error: got = %v", got)
		}
	})
}
