package xptr

func Val[T any](ptr *T) T {
	var zero T

	if ptr == nil {
		return zero
	}

	return *ptr
}

func Ptr[T any](value T) *T {
	return &value
}
