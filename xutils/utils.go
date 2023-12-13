package xutils

func IFElse[T any](cond bool, trueV T, falseV T) T {
	if cond {
		return trueV
	}
	return falseV
}
