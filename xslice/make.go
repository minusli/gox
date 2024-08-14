package xslice

func Make(vals ...int) []int {
	start, end, step := 0, 0, 1
	if len(vals) == 1 {
		end = vals[0]
	}

	if len(vals) == 2 {
		start, end = vals[0], vals[1]
	}

	if len(vals) == 3 {
		start, end, step = vals[0], vals[1], vals[2]
	}

	if len(vals) >= 4 {
		return nil
	}

	if step <= 0 {
		return nil
	}

	if start >= end {
		return nil
	}

	ret := make([]int, 0, end-start)
	for i := start; i < end; i += step {
		ret = append(ret, i)
	}

	return ret
}
