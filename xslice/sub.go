package xslice

func SubSlice[T any](items []T, start, end, step int) []T {
	total := len(items)
	if start < 0 {
		start = 0
	}
	if start > total {
		start = total
	}
	if end < 0 {
		end = total
	}
	if end > total {
		end = total
	}
	if step <= 0 {
		step = 1
	}
	if start >= end {
		return nil
	}

	var ret []T
	for ; start < end; start += step {
		ret = append(ret, items[start])
	}

	return ret
}
