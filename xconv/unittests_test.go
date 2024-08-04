package xconv

import (
	"testing"
)

func TestConv(t *testing.T) {
	t.Run("Conv String To Number", func(t *testing.T) {
		if got := ConvStringToNumber("1", 0); got != 1 { // 整数 -> 整数
			t.Errorf("unittest error: got = %v", got)
		}
		if got := ConvStringToNumber("1.1", 0); got != 1 { // 小数 -> 整数
			t.Errorf("unittest error: got = %v", got)
		}
		if got := ConvStringToNumber("1.1", 0.0); got != 1.1 { // 小数 -> 小数
			t.Errorf("unittest error: got = %v", got)
		}
		if got := ConvStringToNumber("abc", 1.1); got != 1.1 { // 错误
			t.Errorf("unittest error: got = %v", got)
		}
	})

	t.Run("Conv Number To String", func(t *testing.T) {
		if got := ConvNumberToString(1); got != "1" { // 整数 -> string
			t.Errorf("unittest error: got = %v", got)
		}
		if got := ConvNumberToString(1.1); got != "1.1" { // 小数 -> string
			t.Errorf("unittest error: got = %v", got)
		}
	})
}
