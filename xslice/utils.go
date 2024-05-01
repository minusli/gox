package xslice

func ifelse[T any](cond bool, trueV, falseV T) T {
	if cond {
		return trueV
	}

	return falseV
}
