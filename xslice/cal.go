package xslice

// IntersectAny 切片交集
func IntersectAny[T any](items []T, others []T, eq func(a, b T) bool) []T {
	var ret []T

	for _, item := range items {
		if ContainsAny(others, item, eq) {
			ret = append(ret, item)
		}
	}

	return ret
}

// Intersect 切片交集，使用 map 加速计算
func Intersect[T comparable](items []T, others []T) []T {
	var ret []T

	m := ToMap(others, func(a T) (T, bool) { return a, true })
	for _, item := range items {
		if m[item] {
			ret = append(ret, item)
		}
	}

	return ret
}

// Union 切片并集
func Union[T any](items []T, others ...[]T) []T {

	for _, other := range others {
		items = append(items, other...)
	}

	return items
}

// DiffAny 切片差集
func DiffAny[T any](items []T, others []T, eq func(a, b T) bool) []T {
	var ret []T

	for _, item := range items {
		if !ContainsAny(others, item, eq) {
			ret = append(ret, item)
		}
	}

	return ret
}

// Diff 切片交集，使用 map 加速计算
func Diff[T comparable](items []T, others []T) []T {
	var ret []T

	m := ToMap(others, func(a T) (T, bool) { return a, true })
	for _, item := range items {
		if !m[item] {
			ret = append(ret, item)
		}
	}

	return ret
}
