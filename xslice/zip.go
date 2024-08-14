package xslice

func Zip[K comparable, V any](keys []K, values []V) map[K]V {
	ret := make(map[K]V)
	if len(keys) != len(values) {
		return ret
	}

	for idx, k := range keys {
		ret[k] = values[idx]
	}

	return ret
}
