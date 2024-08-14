package xslice

func Sub[T any](items []T, start, end, step int) []T {
	var ret []T
	if start < 0 {
		start += len(items)
	}
	if end < 0 {
		end += len(items)
	}
	if step == 0 {
		step = 1
	}

	if start < end && step > 0 {
		for start = ifelse(start < 0, 0, start); start < end && start < len(items); start += step {
			ret = append(ret, items[start])
		}
		return ret
	}

	if start > end && step < 0 {
		for start = ifelse(start >= len(items), len(items)-1, start); start > end && start >= 0; start += step {
			ret = append(ret, items[start])
		}
		return ret
	}

	return ret
}

func Chunk[T any](items []T, size int) [][]T {
	var ret [][]T

	if size <= 0 {
		return [][]T{items}
	}

	for idx, item := range items {
		chunkIdx := idx / size

		if chunkIdx >= len(ret) {
			ret = append(ret, []T{})
		}

		ret[chunkIdx] = append(ret[chunkIdx], item)
	}

	return ret
}
