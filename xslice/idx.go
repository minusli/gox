package xslice

func First[T any](items []T) (ret T) {
	if len(items) == 0 {
		return
	}

	return items[0]
}

func Last[T any](items []T) (ret T) {
	if len(items) == 0 {
		return
	}

	return items[len(items)-1]
}

func Mid[T any](items []T) (left T, right T) {
	if len(items) == 0 {
		return
	}
	if len(items)&1 == 1 {
		return items[len(items)/2], items[len(items)/2]
	} else {
		return items[len(items)/2-1], items[len(items)/2]
	}

}
