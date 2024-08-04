package xconv

import (
	"fmt"

	"github.com/minusli/gox/xtype"
)

func ConvNumberToString[T xtype.Number](value T) string {
	return fmt.Sprintf("%v", value)
}

func ConvStringToNumber[T xtype.Number](value string, default_ T) T {
	var result T
	_, err := fmt.Sscanf(value, "%v", &result)
	if err != nil {
		return default_
	}

	return result
}
