package slice

func Filter[T any](haystack []T, fn func(item T, index int64) bool) []T {
	result := make([]T, 0)

	for i, v := range haystack {
		if fn(v, int64(i)) {
			result = append(result, v)
		}
	}

	return result
}
